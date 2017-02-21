museclient
==========

museclient is a simple website that will allow a user to selectively search job
listings via [The Muse's API](https://www.themuse.com/developers/api/v2).

museclient uses a "classic" design (pre CSS :)

# Usage

To run museclient, you'll need to have [Golang](https://golang.org/dl/)
installed.

You may optionally register your app to obtain an API key which will allow you
to make many more requests per hour.

Compile the program using `go build` and run it. You may include any optional
flags, such as `--apikey`.

```
./museclient --apikey "your-api-key"
```

# TODO (e.g. further work given more time)

1. **Add tests**

  With more time, an automated test suite would have been added. The primary
  reason that no test suite currently exists is based on both the fact that
  Golang is a statically typed language, which prevents many smaller errors, but
  also that manual testing of code is a frequent part of API exploration. Were
  the project to get larger than it is now, though, lack of a test suite would
  soon become a much larger issue.

2. **Fix known bugs**

  For example, though `museclient` won't let you try and look at jobs on pages
  before 0 in the api, it will happily let you click `Next Page` until you run
  out of requests.

  It also doesn't differentiate API errors at all.

3. **Separate out templates**

  Presently, because there are only 2 pages, templates are not separated from
  the rest of the code, nor are any part of them reused. If more pages were
  added, it would soon be necessary to come up with some kind of templating
  system to improve the pace of developing new templates.

4. **Improve the interface**

  A web page with literally zero CSS and no attention to design beyond base
  functionality isn't going to be fun for anyone to use. Improving the design
  would definitely be a priority.
