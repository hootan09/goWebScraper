#### Go Web Scraper Sample

**this repo Uses Colly Package**
https://github.com/gocolly/colly

this is sample of Scraping https:\\\divar.ir

Run
```sh
$ go mod tidy
$ go run .
```


##### OutPut
```sh
$ go run .
Visiting https://divar.ir/s/tehran
Got a response from https://divar.ir/s/tehran
Finished https://divar.ir/s/tehran
Writing data to file
Data written to file successfully
```

```json
[
    {
        "Name": "iphone 12pro 256G",
        "Image": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMkAAADJCAQAAAAjz96OAAABLElEQVR42u3RQREAAAjDMObfJzbABccjldCkp/SqIEEiJEiEBImQIBESIUEiJEiEBImQIDEBiZAgERIkQoJESIQEiZAgERIkQoJESIQEiZAgERIkQiIkSIQEiZAgERIkQiIkSIQEiZAgERIhQSIkSIQEiZAgERIhQSIkSIQEiZAICRIhQSIkSIQEiZAICRIhQSIkSIRESJAICRIhQSIkSIRESJAICRIhQSIkQoJESJAICRIhQSIkQoJESJAICRIhERIkQoJESJAICRIhERIkQoJESJAICRIkSIQEiZAgERIkQiIkSIQEiZAgERIkQiIkSIQEiZAgERIhQSIkSIQEiZAgERIhQSIkSIQEiZAICRIhQSIkSIQEiZAICRIhQSIkSIRESJAICRIhQaKjFod2faZjQ1LbAAAAAElFTkSuQmCC",
        "Price": "۱۶,۰۰۰,۰۰۰ تومان",
        "Url": "https://divar.ir/v//v/iphone-12pro-256g_گوشی-موبایل_تهران_صادقیه_دیوار/QYhNYG4p",
        "Location": "فروشگاه ‏rent store در صادقیه"
    },
    {
        "Name": "iphone 12 pro max",
        "Image": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMkAAADJCAQAAAAjz96OAAABLElEQVR42u3RQREAAAjDMObfJzbABccjldCkp/SqIEEiJEiEBImQIBESIUEiJEiEBImQIDEBiZAgERIkQoJESIQEiZAgERIkQoJESIQEiZAgERIkQiIkSIQEiZAgERIkQiIkSIQEiZAgERIhQSIkSIQEiZAgERIhQSIkSIQEiZAICRIhQSIkSIQEiZAICRIhQSIkSIRESJAICRIhQSIkSIRESJAICRIhQSIkQoJESJAICRIhQSIkQoJESJAICRIhERIkQoJESJAICRIhERIkQoJESJAICRIkSIQEiZAgERIkQiIkSIQEiZAgERIkQiIkSIQEiZAgERIhQSIkSIQEiZAgERIhQSIkSIQEiZAICRIhQSIkSIQEiZAICRIhQSIkSIRESJAICRIhQaKjFod2faZjQ1LbAAAAAElFTkSuQmCC",
        "Price": "۱۸,۵۰۰,۰۰۰ تومان",
        "Url": "https://divar.ir/v//v/iphone-12-pro-max_گوشی-موبایل_تهران_صادقیه_دیوار/QYzEyxK2",
        "Location": "فروشگاه ‏rent store در صادقیه"
    },
    {
        "Name": "اپل iPhone 13 با حافظهٔ ۱۲۸ گیگابایت",
        "Image": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMkAAADJCAQAAAAjz96OAAABLElEQVR42u3RQREAAAjDMObfJzbABccjldCkp/SqIEEiJEiEBImQIBESIUEiJEiEBImQIDEBiZAgERIkQoJESIQEiZAgERIkQoJESIQEiZAgERIkQiIkSIQEiZAgERIkQiIkSIQEiZAgERIhQSIkSIQEiZAgERIhQSIkSIQEiZAICRIhQSIkSIQEiZAICRIhQSIkSIRESJAICRIhQSIkSIRESJAICRIhQSIkQoJESJAICRIhQSIkQoJESJAICRIhERIkQoJESJAICRIhERIkQoJESJAICRIkSIQEiZAgERIkQiIkSIQEiZAgERIkQiIkSIQEiZAgERIhQSIkSIQEiZAgERIhQSIkSIQEiZAICRIhQSIkSIQEiZAICRIhQSIkSIRESJAICRIhQaKjFod2faZjQ1LbAAAAAElFTkSuQmCC",
        "Price": "۱۵,۰۰۰,۰۰۰ تومان",
        "Url": "https://divar.ir/v//v/اپل-iphone-13-با-حافظهٔ-۱۲۸-گیگابایت_گوشی-موبایل_تهران_صادقیه_دیوار/QYv6CTF7",
        "Location": "فروشگاه ‏rent store در صادقیه"
    },
    ...
]
````
