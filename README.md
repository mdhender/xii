# xii
My take on 12 Factor

# Usage

## AsBool

	development, err := xii.AsBool("DEVELOPMENT", xii.BoolOpts{DefaultValue: false})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[xii] %-30s == %v\n", "DEVELOPMENT", development)

## AsInt

	port, err := xii.AsInt("PORT", xii.IntOpts{DefaultValue: 8080})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[xii] %-30s == %d\n", "PORT", port)

## AsString
	appRoot, err := xii.AsString("APPROOT", xii.StringOpts{DefaultValue: "./"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[xii] %-30s == %q\n", "APPROOT", appRoot)
