package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

//Instruction struct to mp to db table
type Instruction struct {
	ID          int64  `db:"id" json:"id"`
	EventStatus string `db:"event_status" json:"event_status"`
	EventName   string `db:"event_name" json:"event_name"`
}

var dbmap = initDB()

//GetInstructions used for list
func GetInstructions(c *gin.Context) {
	var instructions []Instruction
	_, err := dbmap.Select(&instructions, "SELECT * FROM instruction")
	if err == nil {
		c.JSON(200, instructions)
	} else {
		c.JSON(404, gin.H{"error": "no instruction(s) into the table"})
	}
}

//GetInstruction used for show
func GetInstruction(c *gin.Context) {
	// c.JSON(200, gin.H{"ok": "GET api/v1/instructions/1"})
	id := c.Params.ByName("id")
	var instruction Instruction
	err := dbmap.SelectOne(&instruction, "SELECT * FROM instruction WHERE id=?", id)
	if err == nil {
		instructionID, _ := strconv.ParseInt(id, 10, 64)
		content := &Instruction{
			ID:          instructionID,
			EventStatus: instruction.EventStatus,
			EventName:   instruction.EventName,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "instruction not found"})
	}
}

//PostInstruction used for creation of record.
func PostInstruction(c *gin.Context) {
	var instruction Instruction
	c.Bind(&instruction)

	if instruction.EventStatus != "" && instruction.EventName != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO instruction (event_status, event_name) VALUES (?, ?)`, instruction.EventStatus, instruction.EventName); insert != nil {
			instructionID, err := insert.LastInsertId()
			if err == nil {
				content := &Instruction{
					ID:          instructionID,
					EventStatus: instruction.EventStatus,
					EventName:   instruction.EventName,
				}
				c.JSON(200, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}

	// c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})

}

//UpdateInstruction to update existing record
func UpdateInstruction(c *gin.Context) {
	id := c.Params.ByName("id")
	var instruction Instruction
	err := dbmap.SelectOne(&instruction, "SELECT * FROM instruction WHERE id=?", id)

	if err == nil {
		var json Instruction
		c.Bind(&json)
		instructionID, _ := strconv.ParseInt(id, 0, 64)
		instruction := Instruction{
			ID:          instructionID,
			EventStatus: json.EventStatus,
			EventName:   json.EventName,
		}

		if instruction.EventStatus != "" && instruction.EventName != "" {
			_, err = dbmap.Update(&instruction)

			if err == nil {
				c.JSON(200, instruction)
			} else {
				checkErr(err, "Updated failed")
			}
		} else {
			c.JSON(422, gin.H{"error": "fields are empty"})
		}
	} else {
		c.JSON(404, gin.H{"error": "instruction not found"})
	}
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"event_status\": \"83\", \"event_name\": \"100\" }" http://localhost:8080/api/v1/instructions/1
}

//DeleteInstruction to delete a record
func DeleteInstruction(c *gin.Context) {
	id := c.Params.ByName("id")
	var instruction Instruction
	err := dbmap.SelectOne(&instruction, "SELECT id FROM Instruction WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&instruction)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: " deleted"})
		} else {
			checkErr(err, "Delete failed")
		}
	} else {
		c.JSON(404, gin.H{"error": "instruction not found"})
	}
	// curl -X "DELETE" http://localhost:8080/api/v1/instructions/1
}
