### ISO-3166 Country and Dependent Territories Lists with UN Regional Codes

These lists are the result of merging data from two sources, the Wikipedia [ISO 3166-1 article](http://en.wikipedia.org/wiki/ISO_3166-1#Officially_assigned_code_elements) for alpha and numeric country codes, and the [UN Statistics](https://unstats.un.org/unsd/methodology/m49/overview) site for countries' regional, and sub-regional codes. In addition to countries, it includes dependent territories.

The [International Organization for Standardization (ISO)](https://www.iso.org/iso-3166-country-codes.html) site provides partial data (capitalised and sometimes stripped of non-latin ornamentation), but sells the complete data set as a Microsoft Access 2003 database. Other sites give you the numeric and character codes, but there appeared to be no sites that included the associated UN-maintained regional codes in their data sets. I scraped data from the above two websites that is all publicly available already to produce some ready-to-use complete data sets that will hopefully save someone some time who had similar needs.

### What's available?

The data is available in

* JSON
* XML
* CSV

3 versions exist for each format

* `all.format` - Everything I can find, including regional and sub-regional codes
* `slim-2.format` - English name, numeric country code and alpha-2 code (e.g., NZ)
* `slim-3.format` - English name, numeric country code and alpha-3 code (e.g., NZL)

### What does it look like?

Take a peek inside the `all`, `slim-2` and `slim-3` directories for the full lists of JSON, XML and CSV.

Using JSON as an example:

#### all.json

    [
      {
        "name":"Nigeria",
        "alpha-2":"NG",
        "alpha-3":"NGA",
        "country-code":"566",
        "iso_3166-2":"ISO 3166-2:NG",
        "region":"Africa",
        "sub-region":"Sub-Saharan Africa",
        "intermediate-region":"Western Africa",
        "region-code":"002",
        "sub-region-code":"202",
        "intermediate-region-code":"011"
      },
      // ...
    ]

#### slim-2.json

    [
      {
        "name":"New Zealand",
        "alpha-2":"NZ",
        "country-code":"554"
      },
      // ...
    ]

#### slim-3.json

    [
      {
        "name":"New Zealand",
        "alpha-3":"NZL",
        "country-code":"554"
      },
      // ...
    ]

### Timestamp

* UN Statistical data retrieved 8 December 2020
* Wikipedia data retrieved 8 December 2020, from a document last revised 19 November 2020