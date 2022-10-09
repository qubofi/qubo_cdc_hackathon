package security

import (
	"qubo/qubo-backend/infrastructure/redis"
	"strconv"
	"time"
)

func CreateAuth(userId uint64, tokenDetail *TokenDetails) error {
	authToken := time.Unix(tokenDetail.AtExpires, 0) //converting Unix to UTC(to Time object)
	refreshToken := time.Unix(tokenDetail.RtExpires, 0)
	now := time.Now()

	errAccess := redis.Client.Set(tokenDetail.AccessUuid, userId, authToken.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := redis.Client.Set(tokenDetail.RefreshUuid, userId, refreshToken.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func FetchAuth(authD *AccessDetails) (uint64, error) {
	userid, err := redis.Client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := redis.Client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
