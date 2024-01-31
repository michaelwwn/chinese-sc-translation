# 跨平台繁簡轉換工具

## Overview
This Go program provides a convenient way to translate text between Simplified Chinese and Traditional Chinese using a keyboard-triggered conversation system. The application runs as a daemon and utilizes system tray integration for easy access and control. It listens for a specific hotkey combination, captures the selected text, translates it, and then pastes the translated text back.

## Features
- **Daemon-based Architecture**: Runs in the background as a daemon process.
- **System Tray Integration**: Easily accessible through the system tray icon.
- **Hotkey Triggered Translation**: Translates text using a custom hotkey combination.
- **Automatic Text Capture and Paste**: Captures selected text and pastes the translated text back into the application.
- **Log File Generation**: Generates log files for tracking and debugging.
- **Simplified and Traditional Chinese Translation**: Supports translation between Simplified Chinese (zh-CN) and Traditional Chinese (zh-TW).

## Usage
1. Start the program. It will run as a daemon and appear in the system tray.
2. Select the text you want to translate in any application.
3. Press the configured hotkey (default: Ctrl+Shift+S) to trigger the translation.
4. The program captures the selected text, translates it, and pastes the translated text back.

## Configuration
- **Hotkey**: Default set to Ctrl+Shift+S. Modify the `listenHotkey` function to change the hotkey.
- **Translation Direction**: Default set to convert from Simplified to Traditional Chinese. Modify the `gocc.New("hk2s")` line in `listenHotkey` function for different translation modes.

## Dependencies
- [systray](https://github.com/getlantern/systray): For system tray integration.
- [robotgo](https://github.com/go-vgo/robotgo): For simulating keyboard and mouse actions.
- [gocc](https://github.com/liuzl/gocc): For Chinese character conversion.
- [go-daemon](https://github.com/sevlyar/go-daemon): For daemon functionality.
- [hotkey](https://golang.design/x/hotkey): For hotkey management.

## Contributing
Contributions to enhance the functionality or fix issues in the program are welcome. Please submit pull requests or open issues on the GitHub repository.

## License
This software is released under the [MIT License](https://opensource.org/licenses/MIT).
