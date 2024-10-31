{
  _images+:: {
    logstore: 'acme/logstore:2.9.2',

    read: self.logstore,
    write: self.logstore,
  },
}
