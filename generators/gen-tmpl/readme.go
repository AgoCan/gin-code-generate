package tmpl

// ReadmeContent readme
var ReadmeContent = `# {{ .ProjectName }}
项目启动
` + "\n```\n" + "go build -o app .\n ./app -c {config_file_path} \n" + "```\n" +
	"\n```\n" + "docker build -t test:v-app .\ndocker run -it --rm -p 9000:9000 test:v-app\n" + "```\n"
