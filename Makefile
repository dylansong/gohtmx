.PHONY: dev air tailwind templ browser-sync

# 主命令,启动所有开发服务
dev:
	@echo "启动开发服务..."
	@osascript -e 'tell application "iTerm2"' \
		-e 'set newWindow to (create window with default profile)' \
		-e 'tell newWindow' \
		-e '    tell current session' \
		-e '        write text "cd $(PWD) && air"' \
		-e '        set newPane to (split vertically with default profile)' \
		-e '        tell newPane' \
		-e '            write text "cd $(PWD) && npx tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --watch --minify"' \
		-e '        end tell' \
		-e '        set newPane2 to (split horizontally with default profile)' \
		-e '        tell newPane2' \
		-e '            write text "cd $(PWD) && templ generate --watch"' \
		-e '        end tell' \
		-e '        tell newPane' \
		-e '            set newPane3 to (split horizontally with default profile)' \
		-e '            tell newPane3' \
		-e '                write text "cd $(PWD) && npx browser-sync start --proxy \"localhost:8080\" --files \"**/*.templ,**/*.css, **/*.js, **/*.go\""' \
		-e '            end tell' \
		-e '        end tell' \
		-e '    end tell' \
		-e 'end tell' \
		-e 'end tell'
