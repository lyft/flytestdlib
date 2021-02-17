// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package storage

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeRaw_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_type", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("type", testValue)
			if vString, err := cmdFlags.GetString("type"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Type)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_connection.endpoint", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.Connection.Endpoint.String()

			cmdFlags.Set("connection.endpoint", testValue)
			if vString, err := cmdFlags.GetString("connection.endpoint"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Connection.Endpoint)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_connection.auth-type", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("connection.auth-type", testValue)
			if vString, err := cmdFlags.GetString("connection.auth-type"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Connection.AuthType)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_connection.access-key", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("connection.access-key", testValue)
			if vString, err := cmdFlags.GetString("connection.access-key"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Connection.AccessKey)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_connection.secret-key", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("connection.secret-key", testValue)
			if vString, err := cmdFlags.GetString("connection.secret-key"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Connection.SecretKey)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_connection.region", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("connection.region", testValue)
			if vString, err := cmdFlags.GetString("connection.region"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Connection.Region)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_connection.disable-ssl", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("connection.disable-ssl", testValue)
			if vBool, err := cmdFlags.GetBool("connection.disable-ssl"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.Connection.DisableSSL)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_container", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("container", testValue)
			if vString, err := cmdFlags.GetString("container"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.InitContainer)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_enable-multicontainer", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("enable-multicontainer", testValue)
			if vBool, err := cmdFlags.GetBool("enable-multicontainer"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.MultiContainerEnabled)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_cache.max_size_mbs", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("cache.max_size_mbs", testValue)
			if vInt, err := cmdFlags.GetInt("cache.max_size_mbs"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Cache.MaxSizeMegabytes)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_cache.target_gc_percent", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("cache.target_gc_percent", testValue)
			if vInt, err := cmdFlags.GetInt("cache.target_gc_percent"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Cache.TargetGCPercent)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_limits.maxDownloadMBs", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("limits.maxDownloadMBs", testValue)
			if vInt64, err := cmdFlags.GetInt64("limits.maxDownloadMBs"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt64), &actual.Limits.GetLimitMegabytes)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_defaultHttpClient.timeout", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.DefaultHTTPClient.Timeout.String()

			cmdFlags.Set("defaultHttpClient.timeout", testValue)
			if vString, err := cmdFlags.GetString("defaultHttpClient.timeout"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.DefaultHTTPClient.Timeout)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
