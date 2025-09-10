# encoding: utf-8
require "logstash/devutils/rspec/spec_helper"
require "logstash/outputs/logstore"
require "logstash/codecs/plain"
require "logstash/event"
require "net/http"
require 'webmock/rspec'
include Logstore

describe LogStash::Outputs::Logstore do

  let (:simple_logstore_config) { {'url' => 'http://localhost:3100'} }

  context 'when initializing' do
    it "should register" do
      logstore = LogStash::Plugin.lookup("output", "logstore").new(simple_logstore_config)
      expect { logstore.register }.to_not raise_error
    end

    it 'should populate logstore config with default or initialized values' do
      logstore = LogStash::Outputs::Logstore.new(simple_logstore_config)
      expect(logstore.url).to eql 'http://localhost:3100'
      expect(logstore.tenant_id).to eql nil
      expect(logstore.batch_size).to eql 102400
      expect(logstore.batch_wait).to eql 1
    end
  end

  context 'when adding en entry to the batch' do
    let (:simple_logstore_config) {{'url' => 'http://localhost:3100'}}
    let (:entry) {Entry.new(LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","@timestamp"=>Time.at(1)}),"message", [], [])}
    let (:lbs) {{"buzz"=>"bar","cluster"=>"us-central1"}.sort.to_h}
    let (:include_logstore_config) {{ 'url' => 'http://localhost:3100', 'include_fields' => ["cluster"] }}
    let (:include_entry) {Entry.new(LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","@timestamp"=>Time.at(1)}),"message", ["cluster"], [])}
    let (:include_lbs) {{"cluster"=>"us-central1"}.sort.to_h}

    it 'should not add empty line' do
      plugin = LogStash::Plugin.lookup("output", "logstore").new(simple_logstore_config)
      emptyEntry = Entry.new(LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","@timestamp"=>Time.at(1)}),"foo", [], [])
      expect(plugin.add_entry_to_batch(emptyEntry)).to eql true
      expect(plugin.batch).to eql nil
    end

    it 'should add entry' do
      plugin = LogStash::Plugin.lookup("output", "logstore").new(simple_logstore_config)
      expect(plugin.batch).to eql nil
      expect(plugin.add_entry_to_batch(entry)).to eql true
      expect(plugin.add_entry_to_batch(entry)).to eql true
      expect(plugin.batch).not_to be_nil
      expect(plugin.batch.streams.length).to eq 1
      expect(plugin.batch.streams[lbs.to_s]['entries'].length).to eq 2
      expect(plugin.batch.streams[lbs.to_s]['labels']).to eq lbs
      expect(plugin.batch.size_bytes).to eq 14
    end

    it 'should only allowed labels defined in include_fields' do
      plugin = LogStash::Plugin.lookup("output", "logstore").new(include_logstore_config)
      expect(plugin.batch).to eql nil
      expect(plugin.add_entry_to_batch(include_entry)).to eql true
      expect(plugin.add_entry_to_batch(include_entry)).to eql true
      expect(plugin.batch).not_to be_nil
      expect(plugin.batch.streams.length).to eq 1
      expect(plugin.batch.streams[include_lbs.to_s]['entries'].length).to eq 2
      expect(plugin.batch.streams[include_lbs.to_s]['labels']).to eq include_lbs
      expect(plugin.batch.size_bytes).to eq 14
    end

    it 'should not add if full' do
      plugin = LogStash::Plugin.lookup("output", "logstore").new(simple_logstore_config.merge!({'batch_size'=>10}))
      expect(plugin.batch).to eql nil
      expect(plugin.add_entry_to_batch(entry)).to eql true # first entry is fine.
      expect(plugin.batch).not_to be_nil
      expect(plugin.batch.streams.length).to eq 1
      expect(plugin.batch.streams[lbs.to_s]['entries'].length).to eq 1
      expect(plugin.batch.streams[lbs.to_s]['labels']).to eq lbs
      expect(plugin.batch.size_bytes).to eq 7
      expect(plugin.add_entry_to_batch(entry)).to eql false # second entry goes over the limit.
      expect(plugin.batch).not_to be_nil
      expect(plugin.batch.streams.length).to eq 1
      expect(plugin.batch.streams[lbs.to_s]['entries'].length).to eq 1
      expect(plugin.batch.streams[lbs.to_s]['labels']).to eq lbs
      expect(plugin.batch.size_bytes).to eq 7
    end
  end

  context 'when building json from batch to send' do
    let (:basic_logstore_config) {{'url' => 'http://localhost:3100'}}
    let (:basic_entry) {Entry.new(LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","trace_id"=>"trace_001","@timestamp"=>Time.at(1)}),"message", [], [])}
    let (:include_logstore_config) {{ 'url' => 'http://localhost:3100', 'include_fields' => ["cluster"] }}
    let (:include_entry) {Entry.new(LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","trace_id"=>"trace_001","@timestamp"=>Time.at(1)}),"message", ["cluster"], [])}
    let (:metadata_logstore_config) {{ 'url' => 'http://localhost:3100', 'include_fields' => ["cluster"], 'metadata_fields' => ["trace_id"] }}
    let (:metadata_entry) {Entry.new(LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","trace_id"=>"trace_001","@timestamp"=>Time.at(1)}),"message", ["cluster"], ["trace_id"])}
    let (:metadata_multi_logstore_config) {{ 'url' => 'http://localhost:3100', 'include_fields' => ["cluster"], 'metadata_fields' => ["trace_id", "user_id"] }}
    let (:metadata_multi_entry) {Entry.new(LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","trace_id"=>"trace_001","user_id"=>"user_001","@timestamp"=>Time.at(1)}),"message", ["cluster"], ["trace_id", "user_id"])}

    it 'should not include labels or metadata' do
      plugin = LogStash::Plugin.lookup("output", "logstore").new(basic_logstore_config)
      expect(plugin.batch).to eql nil
      expect(plugin.add_entry_to_batch(basic_entry)).to eql true
      expect(plugin.batch).not_to be_nil
      expect(plugin.batch.to_json).to eq '{"streams":[{"stream":{"buzz":"bar","cluster":"us-central1","trace_id":"trace_001"},"values":[["1000000000","foobuzz"]]}]}'
    end

    it 'should include metadata with no labels' do
      plugin = LogStash::Plugin.lookup("output", "logstore").new(metadata_logstore_config)
      expect(plugin.batch).to eql nil
      expect(plugin.add_entry_to_batch(metadata_entry)).to eql true
      expect(plugin.batch).not_to be_nil
      expect(plugin.batch.to_json).to eq '{"streams":[{"stream":{"cluster":"us-central1"},"values":[["1000000000","foobuzz",{"trace_id":"trace_001"}]]}]}'
    end

    it 'should include labels with no metadata' do
      plugin = LogStash::Plugin.lookup("output", "logstore").new(include_logstore_config)
      expect(plugin.batch).to eql nil
      expect(plugin.add_entry_to_batch(include_entry)).to eql true
      expect(plugin.batch).not_to be_nil
      expect(plugin.batch.to_json).to eq '{"streams":[{"stream":{"cluster":"us-central1"},"values":[["1000000000","foobuzz"]]}]}'
    end

    it 'should include labels with multiple metadata' do
      plugin = LogStash::Plugin.lookup("output", "logstore").new(metadata_multi_logstore_config)
      expect(plugin.batch).to eql nil
      expect(plugin.add_entry_to_batch(metadata_multi_entry)).to eql true
      expect(plugin.batch).not_to be_nil
      expect(plugin.batch.to_json).to eq '{"streams":[{"stream":{"cluster":"us-central1"},"values":[["1000000000","foobuzz",{"trace_id":"trace_001","user_id":"user_001"}]]}]}'
    end
  end

  context 'batch expiration' do
    let (:entry) {Entry.new(LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","@timestamp"=>Time.at(1)}),"message", [], [])}

    it 'should not expire if empty' do
      logstore = LogStash::Outputs::Logstore.new(simple_logstore_config.merge!({'batch_wait'=>0.5}))
      sleep(1)
      expect(logstore.is_batch_expired).to be false
    end
    it 'should not expire batch if not old' do
      logstore = LogStash::Outputs::Logstore.new(simple_logstore_config.merge!({'batch_wait'=>0.5}))
      expect(logstore.add_entry_to_batch(entry)).to eql true
      expect(logstore.is_batch_expired).to be false
    end
    it 'should expire if old' do
      logstore = LogStash::Outputs::Logstore.new(simple_logstore_config.merge!({'batch_wait'=>0.5}))
      expect(logstore.add_entry_to_batch(entry)).to eql true
      sleep(1)
      expect(logstore.is_batch_expired).to be true
    end
  end

  context 'channel' do
    let (:event) {LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","@timestamp"=>Time.at(1)})}

    it 'should send entry if batch size reached with no tenant' do
      logstore = LogStash::Outputs::Logstore.new(simple_logstore_config.merge!({'batch_wait'=>0.5,'batch_size'=>10}))
      logstore.register
      sent = Queue.new
      allow(logstore).to receive(:send) do |batch|
        Thread.new do
          sent << batch
        end
      end
      logstore.receive(event)
      logstore.receive(event)
      sent.deq
      sent.deq
      logstore.close
    end
    it 'should send entry while closing' do
      logstore = LogStash::Outputs::Logstore.new(simple_logstore_config.merge!({'batch_wait'=>10,'batch_size'=>10}))
      logstore.register
      sent = Queue.new
      allow(logstore).to receive(:send) do | batch|
        Thread.new  do
          sent << batch
        end
      end
      logstore.receive(event)
      logstore.close
      sent.deq
    end
    it 'should send entry when batch is expiring' do
      logstore = LogStash::Outputs::Logstore.new(simple_logstore_config.merge!({'batch_wait'=>0.5,'batch_size'=>10}))
      logstore.register
      sent = Queue.new
      allow(logstore).to receive(:send) do | batch|
        Thread.new  do
          sent << batch
        end
      end
      logstore.receive(event)
      sent.deq
      sleep(0.01) # Adding a minimal sleep. In few cases @batch=nil might happen after evaluating for nil
      expect(logstore.batch).to be_nil
      logstore.close
    end
  end

  context 'http requests' do
    let (:entry) {Entry.new(LogStash::Event.new({"message"=>"foobuzz","buzz"=>"bar","cluster"=>"us-central1","@timestamp"=>Time.at(1)}),"message", [], [])}

    it 'should send credentials' do
      conf = {
          'url'=>'http://localhost:3100/logstore/api/v1/push',
          'username' => 'foo',
          'password' => 'bar',
          'tenant_id' => 't'
      }
      logstore = LogStash::Outputs::Logstore.new(conf)
      logstore.register
      b = Batch.new(entry)
      post = stub_request(:post, "http://localhost:3100/logstore/api/v1/push").with(
          basic_auth: ['foo', 'bar'],
          body: b.to_json,
          headers:{
              'Content-Type' => 'application/json' ,
              'User-Agent' => 'logstore-logstash',
              'X-Scope-OrgID'=>'t',
              'Accept'=>'*/*',
              'Accept-Encoding'=>'gzip;q=1.0,deflate;q=0.6,identity;q=0.3',
          }
      )
      logstore.send(b)
      expect(post).to have_been_requested.times(1)
    end

    it 'should not send credentials' do
      conf = {
          'url'=>'http://foo.com/logstore/api/v1/push',
      }
      logstore = LogStash::Outputs::Logstore.new(conf)
      logstore.register
      b = Batch.new(entry)
      post = stub_request(:post, "http://foo.com/logstore/api/v1/push").with(
          body: b.to_json,
          headers:{
              'Content-Type' => 'application/json' ,
              'User-Agent' => 'logstore-logstash',
              'Accept'=>'*/*',
              'Accept-Encoding'=>'gzip;q=1.0,deflate;q=0.6,identity;q=0.3',
          }
      )
      logstore.send(b)
      expect(post).to have_been_requested.times(1)
    end
    it 'should retry 500' do
      conf = {
          'url'=>'http://foo.com/logstore/api/v1/push',
          'retries' => 3,
      }
      logstore = LogStash::Outputs::Logstore.new(conf)
      logstore.register
      b = Batch.new(entry)
      post = stub_request(:post, "http://foo.com/logstore/api/v1/push").with(
          body: b.to_json,
      ).to_return(status: [500, "Internal Server Error"])
      logstore.send(b)
      logstore.close
      expect(post).to have_been_requested.times(3)
    end
    it 'should retry 429' do
      conf = {
          'url'=>'http://foo.com/logstore/api/v1/push',
          'retries' => 2,
      }
      logstore = LogStash::Outputs::Logstore.new(conf)
      logstore.register
      b = Batch.new(entry)
      post = stub_request(:post, "http://foo.com/logstore/api/v1/push").with(
          body: b.to_json,
          ).to_return(status: [429, "stop spamming"])
      logstore.send(b)
      logstore.close
      expect(post).to have_been_requested.times(2)
    end
    it 'should not retry 400' do
      conf = {
          'url'=>'http://foo.com/logstore/api/v1/push',
          'retries' => 11,
      }
      logstore = LogStash::Outputs::Logstore.new(conf)
      logstore.register
      b = Batch.new(entry)
      post = stub_request(:post, "http://foo.com/logstore/api/v1/push").with(
          body: b.to_json,
          ).to_return(status: [400, "bad request"])
      logstore.send(b)
      logstore.close
      expect(post).to have_been_requested.times(1)
    end
    it 'should retry exception' do
      conf = {
          'url'=>'http://foo.com/logstore/api/v1/push',
          'retries' => 11,
      }
      logstore = LogStash::Outputs::Logstore.new(conf)
      logstore.register
      b = Batch.new(entry)
      post = stub_request(:post, "http://foo.com/logstore/api/v1/push").with(
          body: b.to_json,
          ).to_raise("some error").then.to_return(status: [200, "fine !"])
      logstore.send(b)
      logstore.close
      expect(post).to have_been_requested.times(2)
    end
  end
end
