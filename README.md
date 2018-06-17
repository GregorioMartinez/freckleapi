# freckleapi
An incomplete API client for v2 of [Freckle](https://letsfreckle.com/)

## Example Usage
`````

    entrySrv := freckle.NewEntryService(client)
	entries, err := entrySrv.List().Page(0).Enabled(true).Do()
	if err != nil && err != freckle.ErrNoMorePages {
		log.Println(err.Error())
	}

	for _, v := range *entries {
		log.Println(v.Description)			
	}
````