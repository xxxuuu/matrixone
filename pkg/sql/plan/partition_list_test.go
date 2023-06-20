// Copyright 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plan

import "testing"

// -----------------------List Partition--------------------------------------
func TestListPartition(t *testing.T) {
	sqls := []string{
		`CREATE TABLE client_firms (
			id   INT,
			name VARCHAR(35)
		)
		PARTITION BY LIST (id) (
			PARTITION r0 VALUES IN (1, 5, 9, 13, 17, 21),
			PARTITION r1 VALUES IN (2, 6, 10, 14, 18, 22),
			PARTITION r2 VALUES IN (3, 7, 11, 15, 19, 23),
			PARTITION r3 VALUES IN (4, 8, 12, 16, 20, 24)
		);`,

		`CREATE TABLE employees (
			id INT NOT NULL,
			fname VARCHAR(30),
			lname VARCHAR(30),
			hired DATE NOT NULL DEFAULT '1970-01-01',
			separated DATE NOT NULL DEFAULT '9999-12-31',
			job_code INT,
			store_id INT
		)
		PARTITION BY LIST(store_id) (
			PARTITION pNorth VALUES IN (3,5,6,9,17),
			PARTITION pEast VALUES IN (1,2,10,11,19,20),
			PARTITION pWest VALUES IN (4,12,13,14,18),
			PARTITION pCentral VALUES IN (7,8,15,16)
		);`,

		`CREATE TABLE t1 (
			id   INT PRIMARY KEY,
			name VARCHAR(35)
		)
		PARTITION BY LIST (id) (
			PARTITION r0 VALUES IN (1, 5, 9, 13, 17, 21),
			PARTITION r1 VALUES IN (2, 6, 10, 14, 18, 22),
			PARTITION r2 VALUES IN (3, 7, 11, 15, 19, 23),
			PARTITION r3 VALUES IN (4, 8, 12, 16, 20, 24)
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST (a) (
			PARTITION p0 VALUES IN(0,NULL),
			PARTITION p1 VALUES IN( 1,2 ),
			PARTITION p2 VALUES IN( 3,4 ),
			PARTITION p3 VALUES IN( 5,6 )
		);`,
	}

	mock := NewMockOptimizer(false)
	for _, sql := range sqls {
		t.Log(sql)
		logicPlan, err := buildSingleStmt(mock, t, sql)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		outPutPlan(logicPlan, true, t)
	}
}

func TestListPartitionError(t *testing.T) {
	sqls := []string{
		`CREATE TABLE t1 (
			id   INT,
			name VARCHAR(35)
		)
		PARTITION BY LIST (id);`,

		`CREATE TABLE t2 (
			id   INT,
			name VARCHAR(35)
		)
         PARTITION BY LIST (id) (
			PARTITION r0 VALUES IN (1, 5, 9, 13, 17, 21),
			PARTITION r1 VALUES IN (2, 6, 10, 14, 18, 22),
			PARTITION r2 VALUES IN (3, 7, 11, 15, 19, 23),
			PARTITION r2 VALUES IN (4, 8, 12, 16, 20, 24)
		);`,

		`CREATE TABLE t1 (
			id   INT PRIMARY KEY,
			name VARCHAR(35),
			age INT unsigned
		)
		PARTITION BY LIST (age) (
			PARTITION r0 VALUES IN (1, 5, 9, 13, 17, 21),
			PARTITION r1 VALUES IN (2, 6, 10, 14, 18, 22),
			PARTITION r2 VALUES IN (3, 7, 11, 15, 19, 23),
			PARTITION r3 VALUES IN (4, 8, 12, 16, 20, 24)
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST (a) (
			PARTITION p0 VALUES IN(NULL,NULL),
			PARTITION p1 VALUES IN( 1,2 ),
			PARTITION p2 VALUES IN( 3,4 ),
			PARTITION p3 VALUES IN( 5,6 )
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST (a) (
			PARTITION p0 VALUES IN(NULL,NULL),
			PARTITION p1 VALUES IN( 1,2 ),
			PARTITION p2 VALUES IN( 3,1 ),
			PARTITION p3 VALUES IN( 3,3 )
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST (a) (
			PARTITION p0 VALUES IN(0,NULL),
			PARTITION p1 VALUES IN( 1,2 ),
			PARTITION p2 VALUES IN( 3,4 ),
			PARTITION p3 VALUES LESS THAN (50,20)
		);`,

		`create table pt_table_50(
			col1 tinyint,
			col2 smallint,
			col3 int,
			col4 bigint,
			col5 tinyint unsigned,
			col6 smallint unsigned,
			col7 int unsigned,
			col8 bigint unsigned,
			col9 float,
			col10 double,
			col11 varchar(255),
			col12 Date,
			col13 DateTime,
			col14 timestamp,
			col15 bool,
			col16 decimal(5,2),
			col17 text,
			col18 varchar(255),
			col19 varchar(255),
			col20 text,
			primary key(col4,col3,col11)
			) partition by list(col3) (
			PARTITION r0 VALUES IN (1, 5*2, 9, 13, 17-20, 21),
			PARTITION r1 VALUES IN (2, 6, 10, 7, 18, 22),
			PARTITION r2 VALUES IN (3, 7, 11+6, 15, 19, 23),
			PARTITION r3 VALUES IN (4, 8, 12, 16, 20, 24)
			);`,

		`create table pt_table_50(
			col1 tinyint,
			col2 smallint,
			col3 int,
			col4 bigint,
			col5 tinyint unsigned,
			col6 smallint unsigned,
			col7 int unsigned,
			col8 bigint unsigned,
			col9 float,
			col10 double,
			col11 varchar(255),
			col12 Date,
			col13 DateTime,
			col14 timestamp,
			col15 bool,
			col16 decimal(5,2),
			col17 text,
			col18 varchar(255),
			col19 varchar(255),
			col20 text,
			primary key(col4,col3,col11)
			) partition by list(col3) (
			PARTITION r0 VALUES IN (1, 5*2, 9, 13, 17-20, 21),
			PARTITION r1 VALUES IN (2, 6, 10, 14/2, 18, 22),
			PARTITION r2 VALUES IN (3, 7, 11+6, 15, 19, 23),
			PARTITION r3 VALUES IN (4, 8, 12, 16, 20, 24)
			);`,
	}

	mock := NewMockOptimizer(false)
	for _, sql := range sqls {
		_, err := buildSingleStmt(mock, t, sql)
		t.Log(sql)
		t.Log(err)
		if err == nil {
			t.Fatalf("%+v", err)
		}
	}
}

// -----------------------List Columns Partition--------------------------------------
func TestListColumnsPartition(t *testing.T) {
	sqls := []string{
		`CREATE TABLE lc (
				a INT NULL,
				b INT NULL
			)
			PARTITION BY LIST COLUMNS(a,b) (
				PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
				PARTITION p1 VALUES IN( (0,1), (0,2), (0,3), (1,1), (1,2) ),
				PARTITION p2 VALUES IN( (1,0), (2,0), (2,1), (3,0), (3,1) ),
				PARTITION p3 VALUES IN( (1,3), (2,2), (2,3), (3,2), (3,3) )
			);`,

		`CREATE TABLE customers_1 (
			first_name VARCHAR(25),
			last_name VARCHAR(25),
			street_1 VARCHAR(30),
			street_2 VARCHAR(30),
			city VARCHAR(15),
			renewal DATE
		)
			PARTITION BY LIST COLUMNS(city) (
			PARTITION pRegion_1 VALUES IN('Oskarshamn', 'Högsby', 'Mönsterås'),
			PARTITION pRegion_2 VALUES IN('Vimmerby', 'Hultsfred', 'Västervik'),
			PARTITION pRegion_3 VALUES IN('Nässjö', 'Eksjö', 'Vetlanda'),
			PARTITION pRegion_4 VALUES IN('Uppvidinge', 'Alvesta', 'Växjo')
		);`,

		`CREATE TABLE customers_2 (
			first_name VARCHAR(25),
			last_name VARCHAR(25),
			street_1 VARCHAR(30),
			street_2 VARCHAR(30),
			city VARCHAR(15),
			renewal DATE
		)
		PARTITION BY LIST COLUMNS(renewal) (
			PARTITION pWeek_1 VALUES IN('2010-02-01', '2010-02-02', '2010-02-03',
				'2010-02-04', '2010-02-05', '2010-02-06', '2010-02-07'),
			PARTITION pWeek_2 VALUES IN('2010-02-08', '2010-02-09', '2010-02-10',
				'2010-02-11', '2010-02-12', '2010-02-13', '2010-02-14'),
			PARTITION pWeek_3 VALUES IN('2010-02-15', '2010-02-16', '2010-02-17',
				'2010-02-18', '2010-02-19', '2010-02-20', '2010-02-21'),
			PARTITION pWeek_4 VALUES IN('2010-02-22', '2010-02-23', '2010-02-24',
				'2010-02-25', '2010-02-26', '2010-02-27', '2010-02-28')
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(a,b) PARTITIONS 4 (
			PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
			PARTITION p1 VALUES IN( (0,1), (0,2), (0,3), (1,1), (1,2) ),
			PARTITION p2 VALUES IN( (1,0), (2,0), (2,1), (3,0), (3,1) ),
			PARTITION p3 VALUES IN( (1,3), (2,2), (2,3), (3,2), (3,3) )
		);`,

		`CREATE TABLE lc (
				a INT NULL,
				b INT NULL
			)
			PARTITION BY LIST COLUMNS(a,b) (
				PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
				PARTITION p1 VALUES IN( (0,1), (0,2) ),
				PARTITION p2 VALUES IN( (1,0), (2,0) )
			);`,
	}

	mock := NewMockOptimizer(false)
	for _, sql := range sqls {
		t.Log(sql)
		logicPlan, err := buildSingleStmt(mock, t, sql)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		outPutPlan(logicPlan, true, t)
	}
}

func TestListColumnsPartitionError(t *testing.T) {
	sqls := []string{
		`CREATE TABLE t1 (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(a,b);`,

		`CREATE TABLE t2 (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(a,b) (
			PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
			PARTITION p1 VALUES IN( (0,1), (0,2), (0,3), (1,1), (1,2) ),
			PARTITION p2 VALUES IN( (1,0), (2,0), (2,1), (3,0), (3,1) ),
			PARTITION p2 VALUES IN( (1,3), (2,2), (2,3), (3,2), (3,3) )
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(a,b) PARTITIONS 5 (
			PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
			PARTITION p1 VALUES IN( (0,1), (0,2), (0,3), (1,1), (1,2) ),
			PARTITION p2 VALUES IN( (1,0), (2,0), (2,1), (3,0), (3,1) ),
			PARTITION p3 VALUES IN( (1,3), (2,2), (2,3), (3,2), (3,3) )
		);`,
	}

	mock := NewMockOptimizer(false)
	for _, sql := range sqls {
		_, err := buildSingleStmt(mock, t, sql)
		t.Log(sql)
		t.Log(err)
		if err == nil {
			t.Fatalf("%+v", err)
		}
	}
}

func TestListPartitionFunction(t *testing.T) {
	sqls := []string{
		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(a,b) (
			PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
			PARTITION p1 VALUES IN( (0,1), (0,4+2) ),
			PARTITION p2 VALUES IN( (1,0), (2,0) )
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST(a) (
			PARTITION p0 VALUES IN(0, NULL ),
			PARTITION p1 VALUES IN(1, 2),
			PARTITION p2 VALUES IN(3, 4)
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(b) (
			PARTITION p0 VALUES IN( 0,NULL ),
			PARTITION p1 VALUES IN( 1,2 ),
			PARTITION p2 VALUES IN( 3,4 )
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(b) (
			PARTITION p0 VALUES IN( 0,NULL ),
			PARTITION p1 VALUES IN( 1,1+1 ),
			PARTITION p2 VALUES IN( 3,4 )
		);`,
	}

	mock := NewMockOptimizer(false)
	for _, sql := range sqls {
		t.Log(sql)
		logicPlan, err := buildSingleStmt(mock, t, sql)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		outPutPlan(logicPlan, true, t)
	}
}

func TestListPartitionFunctionError(t *testing.T) {
	sqls := []string{
		`create table pt_table_45(
			col1 tinyint,
			col2 smallint,
			col3 int,
			col4 bigint,
			col5 tinyint unsigned,
			col6 smallint unsigned,
			col7 int unsigned,
			col8 bigint unsigned,
			col9 float,
			col10 double,
			col11 varchar(255),
			col12 Date,
			col13 DateTime,
			col14 timestamp,
			col15 bool,
			col16 decimal(5,2),
			col17 text,
			col18 varchar(255),
			col19 varchar(255),
			col20 text,
			primary key(col4,col3,col11))
		partition by list(col3) (
			PARTITION r0 VALUES IN (1, 5*2, 9, 13, 17-20, 21),
			PARTITION r1 VALUES IN (2, 6, 10, 14/2, 18, 22),
			PARTITION r2 VALUES IN (3, 7, 11+6, 15, 19, 23),
			PARTITION r3 VALUES IN (4, 8, 12, 16, 20, 24)
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(a,b) (
			PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
			PARTITION p1 VALUES IN( (0,1), (0,4/2) ),
			PARTITION p2 VALUES IN( (1,0), (2,0) )
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(a,b) (
			PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
			PARTITION p1 VALUES IN( (0,1), (0,4.2) ),
			PARTITION p2 VALUES IN( (1,0), (2,0) )
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(a,b) (
			PARTITION p0 VALUES IN( 0,NULL ),
			PARTITION p1 VALUES IN( 0,1 ),
			PARTITION p2 VALUES IN( 1,0 )
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST(a) (
			PARTITION p0 VALUES IN(0, NULL ),
			PARTITION p1 VALUES IN(1, 4/2),
			PARTITION p2 VALUES IN(3, 4)
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(b) (
			PARTITION p0 VALUES IN( 0,NULL ),
			PARTITION p1 VALUES IN( 1,4/2 ),
			PARTITION p2 VALUES IN( 3,4 )
		);`,

		`CREATE TABLE lc (
			a INT NULL,
			b INT NULL
		)
		PARTITION BY LIST COLUMNS(a,b) (
			PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
			PARTITION p1 VALUES IN( (0,1,3), (0,4,5) ),
			PARTITION p2 VALUES IN( (1,0), (2,0) )
		);`,
	}

	mock := NewMockOptimizer(false)
	for _, sql := range sqls {
		_, err := buildSingleStmt(mock, t, sql)
		t.Log(sql)
		t.Log(err)
		if err == nil {
			t.Fatalf("%+v", err)
		}
	}
}