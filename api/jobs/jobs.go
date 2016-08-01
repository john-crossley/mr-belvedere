package jobs

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"github.com/john-crossley/mr-belvedere/models"
	"time"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	jobs := []models.Job{}

	err := db.C(models.CollectionJob).Find(nil).Sort("-updated_on").All(&jobs)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, jobs)
}

func Find(c *gin.Context) {
	job := models.Job{}
	db := c.MustGet("db").(*mgo.Database)

	err := db.C(models.CollectionJob).Find(bson.M{"_id": bson.ObjectIdHex(c.Params.ByName("id"))}).One(&job)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, job)
}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	job := models.Job{}
	job.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)
	job.UpdatedOn = job.CreatedOn

	err := c.Bind(&job)
	if err != nil {
		c.Error(err)
	}

	err = db.C(models.CollectionJob).Insert(job)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusCreated, gin.H{"error": false})
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	job := models.Job{}
	err := c.Bind(&job)
	if err != nil {
		c.Error(err)
	}

	id := c.Params.ByName("id")
	doc := bson.M{
		"name": job.Name,
		"status": job.Status,
		"author": job.Author,
		"updated_on": time.Now().UnixNano() / int64(time.Millisecond),
	}
	jobDoc, upsertERR := db.C(models.CollectionJob).UpsertId(id, doc)
	if upsertERR != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, jobDoc)
}

func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	id := c.Params.ByName("id")

	if len(id) < 1 {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Please provide the job _id"})
	}

	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.C(models.CollectionJob).Remove(query)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Job " + id + " has been deleted."})
}