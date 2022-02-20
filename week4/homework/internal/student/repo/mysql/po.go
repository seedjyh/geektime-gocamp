package mysql

import (
	"geektime-gocamp/week4/homework/internal/student"
)

type studentPO struct {
	ID        int64  `xorm:"bigint(20) 'id' not null pk autoincr"`
	UID       string `xorm:"char(64) 'uid' not null unique"`
	RealName  string `xorm:"varchar(128) 'real_name' not null"`
	IsDeleted bool   `xorm:"tinyint 'is_deleted' not null default 0"`
}

func (s *studentPO) TableName() string {
	return "student"
}

func (s *studentPO) initFromStudentDO(do *student.StudentDO) *studentPO {
	s.ID = 0
	s.UID = do.UID.String()
	s.RealName = do.RealName.String()
	s.IsDeleted = false
	return s
}
