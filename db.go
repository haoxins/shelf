package shelf

import r "gopkg.in/dancannon/gorethink.v2"

type DB struct {
	Session *r.Session
	Options r.ConnectOpts
}

func Open(opts r.ConnectOpts) (*DB, error) {
	session, err := r.Connect(opts)

	if err != nil {
		return nil, err
	}

	db := DB{
		Session: session,
		Options: opts,
	}

	return &db, nil
}

func (db *DB) Close(opts ...r.CloseOpts) error {
	return db.Session.Close(opts...)
}

func (db *DB) IsConnected() bool {
	return db.Session.IsConnected()
}

func (db *DB) Reconnect(opts ...r.CloseOpts) error {
	return db.Session.Reconnect(opts...)
}
