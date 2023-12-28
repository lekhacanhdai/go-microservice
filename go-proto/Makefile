gen-go:
	rm -rf openapiv2
	rm -rf generated
	@buf mod update
	@buf generate
	@find "openapiv2" -type f -name "*.json" -iname "*apidocs.swagger.json" -exec cp {} "./swagger" \;
	rm -rf openapiv2
