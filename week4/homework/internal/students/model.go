package students

type UID string

func (uid UID) String() string {
	return string(uid)
}

func (uid UID) Valid() bool {
	return len(uid) > 0
}

type RealName string

func (r RealName) String() string {
	return string(r)
}

func (r RealName) Valid() bool {
	return len(r) > 0
}

type StudentDO struct {
	UID      UID
	RealName RealName
}

func (do StudentDO) Valid() bool {
	return do.UID.Valid() &&
		do.RealName.Valid()
}
