package utils

import "github.com/go-rod/rod/lib/devices"

var (
	IPhone11 = devices.Device{
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) EdgiOS/124.0.2478.50 Version/17.0 Mobile/15E148 Safari/604.1",
		AcceptLanguage: "zh-CN",
		Screen: devices.Screen{
			DevicePixelRatio: 3,
			Horizontal: devices.ScreenSize{
				Width:  812,
				Height: 375,
			},
			Vertical: devices.ScreenSize{
				Width:  375,
				Height: 812,
			}},
		Title: "iPhone 11",
	}
)
