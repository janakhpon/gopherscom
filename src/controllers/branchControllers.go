package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func AddCompanyBranch(c *gin.Context) {
	var branchBody models.Branch
	c.BindJSON(&branchBody)

	branch := models.Branch{
		ID:          uuid.New().String(),
		CID:         c.Request.URL.Query().Get("cid"),
		NAME:        branchBody.NAME,
		ADDRESS:     branchBody.ADDRESS,
		ZIPCODE:     branchBody.ZIPCODE,
		CITY:        branchBody.CITY,
		STATE:       branchBody.STATE,
		COUNTRY:     branchBody.COUNTRY,
		LAT:         branchBody.LAT,
		LON:         branchBody.LON,
		FOUNDEDYEAR: branchBody.FOUNDEDYEAR,
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&branch)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}

	cachebranch, err := json.Marshal(branch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set(branch.ID, cachebranch, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &branch,
	})

	return
}

func UpdateCompanyBranch(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var branchBody models.Branch
	c.BindJSON(&branchBody)
	resbranch := &models.Branch{ID: id}

	err := dbConnect.Select(resbranch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	branch := models.Branch{
		ID:          uuid.New().String(),
		CID:         resbranch.CID,
		NAME:        branchBody.NAME,
		ADDRESS:     branchBody.ADDRESS,
		ZIPCODE:     branchBody.ZIPCODE,
		CITY:        branchBody.CITY,
		STATE:       branchBody.STATE,
		COUNTRY:     branchBody.COUNTRY,
		LAT:         branchBody.LAT,
		LON:         branchBody.LON,
		FOUNDEDYEAR: branchBody.FOUNDEDYEAR,
		CREATEDAT:   resbranch.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}

	updateError := dbConnect.Update(&branch)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}

	cachebranch, err := json.Marshal(branch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set(branch.ID, cachebranch, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &branch,
	})
	return
}

func GetCompanyBranches(c *gin.Context) {
	companyid := c.Request.URL.Query().Get("cid")
	branches := &[]models.Branch{}
	val, err := rdbClient.Get(companyid).Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
		return
	}
	err = json.Unmarshal([]byte(val), &branches)
	if branches != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "succeed",
			"data": branches,
		})
		return
	}
	err = dbConnect.Model(branches).Where("cid = ?", companyid).Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": branches,
	})
	return
}

func DeleteCompanyBranch(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var branchBody models.Branch
	c.BindJSON(&branchBody)
	branch := &models.Branch{ID: id}

	err := dbConnect.Delete(branch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	err = rdbClient.Del(id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}

func GetBranch(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	branch := models.Branch{ID: id}
	res, err := rdbClient.Get("branch" + id).Result()

	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
	} else {
		err = json.Unmarshal([]byte(res), &branch)
		if branch.CID != "" {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "succeed",
				"data":   branch,
				"status": "from redis",
			})
			return
		}
	}

	err = dbConnect.Select(branch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	cachebranch, err := json.Marshal(branch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("branch"+branch.ID, cachebranch, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": branch,
	})
	return
}

func ResetBranchCache(c *gin.Context) {
	var branchList []models.Branch
	var branch models.Branch

	keys := rdbClient.Keys("branch*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &branch)
		if branch.ID != "" {
			branchList = append(branchList, branch)
		}
	}

	if len(branchList) != 0 {
		for _, key := range branchList {
			err := rdbClient.Del("branch" + key.ID).Err()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err,
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":    "resetted cache",
			"status": "from redis",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "failed to reset",
	})
	return
}
