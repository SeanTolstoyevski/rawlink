# Raw Link

This simple program makes backlinks queried from ahref ready for Google Search Console. See footnote 1 for more information.  
It can be used to clean up thousands of malicious backlinks.

**Note**: this program was written at the request of a friend of mine. I don't have an expertise for SEO or Google Search Console, and I'm not a pro in these topics.

## Installation and Usage

Install this program:

```bash
go get github.com/SeanTolstoyevski/rawlink
```

```cmd
rawlink --help
```

Output:

```cmd
Usage of rawlink:
  -i string
        Location of the file to be read (default "backlink.txt")
  -o string
        Name of the file to be written (default "fixed_backlinks.txt")
```

## Footnote 1

Imagine you have an e-commerce site. Your malicious competitors may send your page URL to sites that produce obscene/adult content.
This affects your page's rank on Google.

Example of Ahref output:

```txt
domain:adultsitedomainname.extension/x/y/z/post-name
```

The program converts to:

```txt
domain:adultsitedomainname.extension
```

You can then tell Google **"reject these domains"**.

