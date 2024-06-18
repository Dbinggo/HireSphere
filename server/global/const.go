package global

// 所有可导出的常量读取位置
// 命名规范 ：哪个包下（或者用于哪类包）_哪个分类_具体描述_二级描述
const (
	/**************************COMMON*************************/
	COMMON_EMPTY_STRING = ""
	/**************************COMMON*************************/

	/**************************CONFIG*************************/

	CONFIG_FILE_PATH_DEFAULT = "/config.yaml"
	CONFIG_APP_ENV_PRO       = "pro"
	CONFIG_APP_ENV_DEV       = "dev"

	/**************************CONFIG*************************/

	/**************************LOGGER*************************/
	LOGGER_FORMAT_JSON = "json"

	LOGGER_FILE_INFO_NAME  = "info.log"
	LOGGER_FILE_ERROR_NAME = "error.log"

	LOGGER_KEY_CALLER  = "caller"
	LOGGER_KEY_TRACEID = "traceId"
	/**************************LOGGER*************************/

)
