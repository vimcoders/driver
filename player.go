package driver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vimcoders/sql"
)

type Player struct {
	PlayerID int64 `key:"true"`
	Name     string
	Level    int32
	ItemList ItemList
	StepList []int32
}

func (p *Player) TableName() string {
	return ""
}

type Item struct {
	Type  int32
	Count int32
}

type ItemList []Item

func (it ItemList) ToString() (str string) {
	itemStr := make([]string, len(it))

	for i := 0; i < len(it); i++ {
		itemStr[i] = fmt.Sprintf("%v:%v", it[i].Type, it[i].Count)
	}

	return strings.Join(itemStr, ",")
}

func (it ItemList) Convert(str string) sql.Convertor {
	for _, itemStr := range strings.Split(str, ",") {
		info := strings.Split(itemStr, ":")

		if len(info) <= 2 {
			continue
		}

		t, err := strconv.Atoi(info[0])

		if err != nil {
			continue
		}

		count, err := strconv.Atoi(info[1])

		if err != nil {
			continue
		}

		it = append(it, Item{
			Type:  int32(t),
			Count: int32(count),
		})
	}

	return it
}

//CREATE TABLE `player` (
//  `PlayerID` int(11) NOT NULL AUTO_INCREMENT,
//  `Name` varchar(64) NOT NULL DEFAULT '',
//  `Level` int(11) NOT NULL,
//  `ItemList` text NOT NULL,
//  `StepList` text NOT NULL,
//  `CreatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//  PRIMARY KEY (`PlayerID`)
//) ENGINE=InnoDB AUTO_INCREMENT=663806 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC
///*!50100 PARTITION BY LINEAR HASH (PlayerID)
//PARTITIONS 100 */;
