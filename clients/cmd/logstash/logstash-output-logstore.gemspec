Gem::Specification.new do |s|
  s.name    = 'logstash-output-logstore'
  s.version = '1.2.0'
  s.authors = ['Aditya C S','Cyril Tovena']
  s.email   = ['aditya.gnu@gmail.com','cyril.tovena@acme.com']

  s.summary       = 'Output plugin to ship logs to a Acme Logstore server'
  s.description   = 'Output plugin to ship logs to a Acme Logstore server'
  s.homepage      = 'https://example.com/acme/logstore/'
  s.license       = 'Apache-2.0'
  s.require_paths = ["lib"]

  # Files
  s.files = Dir['lib/**/*','spec/**/*','vendor/**/*','*.gemspec','*.md','CONTRIBUTORS','Gemfile']
   # Tests
  s.test_files = s.files.grep(%r{^(test|spec|features)/})

  # Special flag to let us know this is actually a logstash plugin
  s.metadata = { "logstash_plugin" => "true", "logstash_group" => "output" }

  # Gem dependencies
  #
  s.add_runtime_dependency "logstash-core-plugin-api", ">= 1.60", "<= 2.99"
  s.add_runtime_dependency "logstash-codec-plain", "3.1.0"
  s.add_development_dependency 'logstash-devutils', "2.0.2"
end
