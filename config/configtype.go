package config

type Configuration struct {
	AbortController struct {
		GpioPort int `json:"gpio_port"`
	} `json:"abort_controller"`
	LcdScreenController struct {
		Address string `json:"address"`
	} `json:"lcd_screen_controller"`
	RemoteController struct {
		Device string `json:"device"`
	} `json:"remote_controller"`
	MotionController struct {
		Boards struct {
			Pca96851 struct {
				Address             string `json:"address"`
				ReferenceClockSpeed int    `json:"reference_clock_speed"`
				Frequency           int    `json:"frequency"`
			} `json:"pca9685_1"`
			Pca96852 struct {
				Address             string `json:"address"`
				ReferenceClockSpeed int    `json:"reference_clock_speed"`
				Frequency           int    `json:"frequency"`
			} `json:"pca9685_2"`
		} `json:"boards"`
		Servos struct {
			RearShoulderLeft struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"rear_shoulder_left"`
			RearLegLeft struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"rear_leg_left"`
			RearFeetLeft struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"rear_feet_left"`
			RearShoulderRight struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"rear_shoulder_right"`
			RearLegRight struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"rear_leg_right"`
			RearFeetRight struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"rear_feet_right"`
			FrontShoulderLeft struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"front_shoulder_left"`
			FrontLegLeft struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"front_leg_left"`
			FrontFeetLeft struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"front_feet_left"`
			FrontShoulderRight struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"front_shoulder_right"`
			FrontLegRight struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"front_leg_right"`
			FrontFeetRight struct {
				Pca9685   int `json:"pca9685"`
				Channel   int `json:"channel"`
				MinPulse  int `json:"min_pulse"`
				MaxPulse  int `json:"max_pulse"`
				RestAngle int `json:"rest_angle"`
			} `json:"front_feet_right"`
		} `json:"servos"`
	} `json:"motion_controller"`
}
