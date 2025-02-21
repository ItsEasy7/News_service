package reflection

import (
	"Gogogo/internal/service/news"
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

func CollectEndpointInfo(router *gin.Engine, name, baseURL string) news.EndpointInfo {
	endpointInfo := news.EndpointInfo{
		Name: name,
		Settings: map[string]string{
			"baseUrl": baseURL,
		},
		Metadata: news.Metadata{
			RoutingMap: []news.RoutingMapEntry{},
		},
		Actions: make(map[string]news.Action),
	}

	for _, route := range router.Routes() {
		handlerName := getHandlerName(route.HandlerFunc)
		action := handlerName
		pathParams := parsePathParams(route.Path)

		// Добавление в routingMap
		endpointInfo.Metadata.RoutingMap = append(endpointInfo.Metadata.RoutingMap, news.RoutingMapEntry{
			Action: action,
			Route:  route.Path,
			Method: route.Method,
		})

		// Формирование информации о параметрах
		params := news.Params{}

		// PathParams
		if len(pathParams) > 0 {
			pathSchema := &news.ParamSchema{
				Type:  "object",
				Props: make(map[string]news.Prop),
			}
			for _, param := range pathParams {
				pathSchema.Props[param] = news.Prop{
					Type: "string",
				}
			}
			params.PathParams = pathSchema
		}

		// QueryParams и Body
		if route.HandlerFunc != nil {
			handlerType := reflect.TypeOf(route.HandlerFunc)
			if handlerType.NumIn() > 1 {
				for i := 1; i < handlerType.NumIn(); i++ {
					paramType := handlerType.In(i)
					switch paramType.Kind() {
					case reflect.Struct:
						if route.Method == "GET" || route.Method == "DELETE" {
							params.QueryParams = parseStructToSchema(paramType)
						} else if route.Method == "POST" || route.Method == "PUT" || route.Method == "PATCH" {
							params.Body = parseStructToSchema(paramType)
						}
					}
				}
			}
		}

		endpointInfo.Actions[action] = news.Action{
			Params:  params,
			Handler: route.Path,
		}
	}

	return endpointInfo
}

// parseStructToSchema преобразует структуру в ParamSchema
func parseStructToSchema(paramType reflect.Type) *news.ParamSchema {
	schema := &news.ParamSchema{
		Type:  "object",
		Props: make(map[string]news.Prop),
	}

	for i := 0; i < paramType.NumField(); i++ {
		field := paramType.Field(i)
		jsonTag := field.Tag.Get("json")
		bindingTag := field.Tag.Get("binding")

		// Если тег json не указан, используем имя поля в нижнем регистре
		fieldName := strings.ToLower(field.Name)
		if jsonTag != "" {
			fieldName = strings.Split(jsonTag, ",")[0]
		}

		// Определяем, обязательно ли поле
		isOptional := !strings.Contains(bindingTag, "required")

		// Определяем тип поля
		fieldType := field.Type
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
			isOptional = true
		}

		var propType string
		switch fieldType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			propType = "number"
		case reflect.Float32, reflect.Float64:
			propType = "number"
		case reflect.Bool:
			propType = "boolean"
		case reflect.String:
			propType = "string"
		case reflect.Struct:
			propType = "object"
		case reflect.Slice:
			propType = "array"
		default:
			propType = "unknown"
		}

		// Если поле — структура, рекурсивно создаем схему
		if fieldType.Kind() == reflect.Struct {
			schema.Props[fieldName] = news.Prop{
				Type:     propType,
				Optional: isOptional,
				Props:    parseStructToSchema(fieldType).Props,
			}
		} else {
			schema.Props[fieldName] = news.Prop{
				Type:     propType,
				Optional: isOptional,
			}
		}
	}

	return schema
}

// Получение названия action`а
func getHandlerName(handler gin.HandlerFunc) string {
	if handler == nil {
		return ""
	}
	ptr := reflect.ValueOf(handler).Pointer()
	name := runtime.FuncForPC(ptr).Name()
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
}

// Получение названия pathParams
func parsePathParams(path string) []string {
	re := regexp.MustCompile(`:(\w+)`)
	matches := re.FindAllStringSubmatch(path, -1)
	params := make([]string, 0)
	for _, match := range matches {
		if len(match) >= 2 {
			params = append(params, match[1])
		}
	}
	return params
}

func RegisterEndpoints(scheme string) (string, error) {
	reqBody := bytes.NewBuffer([]byte(scheme))
	req, err := http.NewRequest("POST", "http://localhost/v1/registry/services", reqBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("service-token", "")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return "Запрос успешен", nil
	} else {
		return "", fmt.Errorf("Ошибка: статус код %d", resp.StatusCode)
	}
}
