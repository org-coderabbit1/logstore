{
  _images+:: {
    logstore: 'acme/logstore:2.6.1',

    read: self.logstore,
    write: self.logstore,
  },
}
