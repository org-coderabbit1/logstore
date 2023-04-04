{
  _images+:: {
    logstore: 'acme/logstore:2.7.5',

    read: self.logstore,
    write: self.logstore,
  },
}
