package tutorials

import (
	"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{
		"首页": {"首页", "", homeScreen, true},
		"TCP服务端": {"TCP服务端",
			"TCP服务端",
			tcpServerScreen,
			true,
		},
		"TCP客户端": {"TCP客户端",
			"TCP客户端",
			tcpClientScreen,
			true,
		},
		"HTTP服务端": {"HTTP服务端",
			"TCP客户端",
			httpServerScreen,
			true,
		},
		"HTTP客户端": {"HTTP客户端",
			"HTTP客户端",
			welcomeScreen,
			true,
		},
		"JMS消息": {"JMS消息",
			"JMS消息",
			welcomeScreen,
			true,
		},
		"网络扫描": {"网络扫描",
			"网络扫描",
			welcomeScreen,
			true,
		},
		"系统帮助": {"系统帮助",
			"系统帮助",
			welcomeScreen,
			true,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"": {"首页", "TCP服务端", "TCP客户端", "HTTP服务端", "HTTP客户端", "JMS消息", "网络扫描", "系统帮助"},
	}
)
