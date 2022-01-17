package reps

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"mngr/models"
)

type ConfigRepository struct {
	Connection *redis.Client
}

func (r *ConfigRepository) GetMlConfig() (*models.MlConfig, error) {
	var config = &models.MlConfig{}
	conn := r.Connection
	key := "ml_config"
	data, err := conn.Get(context.Background(), key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return config, nil
		} else {
			log.Println("Error getting sources from redis: ", err)
			return nil, err
		}
	}

	err = json.Unmarshal([]byte(data), config)
	if err != nil {
		log.Println("Error deserializing json: ", err)
		return nil, err
	}

	return config, nil
}

func (r *ConfigRepository) SaveMlConfig(config *models.MlConfig) error {
	conn := r.Connection
	key := "ml_config"
	data, err := json.Marshal(config)
	if err != nil {
		log.Println("Error serializing json: ", err)
		return err
	}

	err = conn.Set(context.Background(), key, data, 0).Err()
	if err != nil {
		log.Println("Error saving ml config to redis: ", err)
		return err
	}

	return nil
}

func (r *ConfigRepository) RestoreMlConfig() (*models.MlConfig, error) {
	objJson := `{"device_config":{"device_name":"gokalp-ubuntu","device_services":[1,2,4],"device_type":0},"handler":{"read_service_overlay":false,"save_image_extension":"jpg","save_image_folder_path":"/home/gokalp/Pictures","show_image_caption":false,"show_image_fullscreen":false,"show_image_wait_key":1},"heartbeat":{"interval":5},"jetson":{"model_name":"ssd-mobilenet-v2","threshold":0.1,"white_list":[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90]},"once_detector":{"imagehash_threshold":3,"psnr_threshold":0.2,"ssim_threshold":0.2},"source_hub":{"buffer_size":2,"fps":1,"kill_starter_proc":true,"max_retry":150,"max_retry_in":6},"redis":{"host":"127.0.0.1","port":6379},"torch":{"model_name":"ultralytics/yolov5","model_name_specific":"yolov5x6","threshold":0.1,"white_list":[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79]}}`

	conn := r.Connection
	key := "ml_config"
	err := conn.Set(context.Background(), key, objJson, 0).Err()
	if err != nil {
		log.Println("Error saving default ml config to redis: ", err)
		return nil, err
	}

	var config = &models.MlConfig{}
	err = json.Unmarshal([]byte(objJson), config)
	if err != nil {
		log.Println("Error deserializing json: ", err)
		return nil, err
	}

	return config, nil
}