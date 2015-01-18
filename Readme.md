# gokeen/query

[![Build Status](https://travis-ci.org/gokeen/query.svg?branch=master)][ci]
[![GoDoc](https://godoc.org/gopkg.in/gokeen/query.v1?status.svg)][gd]

  [ci]: https://travis-ci.org/gokeen/query
  [gd]: http://godoc.org/gopkg.in/gokeen/query.v1

Queries package for gokeen

## Install

    go get github.com/gokeen/query

__gopkg.in__

    go get gopkg.in/gokeen/query.v1

## Example

    k := keen.NewClient("aprojectid", func(c *api.Client) {
      c.ReadKey = "areadkey"
    })

    qry := query.CountUnique("awesome-events", "some.property", func(q *query.Query) {
      q.Timeframe = timeframe.Previous.Hour(2)
    })

    var res Result
    err = k.Query(qry, &res);
    if err != nil {
      // handle query error
    }

## API

### Count resource

https://keen.io/docs/api/reference/#count-resource

    func Count(eventCollection string, opts ...queryOpts) Resource

### CountUnique resource

https://keen.io/docs/api/reference/#count-unique-resource

    func CountUnique(eventCollection, propertyTarget string, opts ...queryOpts) Resource

### Minimum resource

https://keen.io/docs/api/reference/#minimum-resource

    func Minimum(eventCollection, propertyTarget string, opts ...queryOpts) Resource

### Maximum resource

https://keen.io/docs/api/reference/#maximum-resource

    func Maximum(eventCollection, propertyTarget string, opts ...queryOpts) Resource

### Average resource

https://keen.io/docs/api/reference/#average-resource

    func Average(eventCollection, propertyTarget string, opts ...queryOpts) Resource

### Median resource

https://keen.io/docs/api/reference/#median-resource

    func Median(eventCollection, propertyTarget string, opts ...queryOpts) Resource

### Percentile resource

https://keen.io/docs/api/reference/#percentile-resource

    func Percentile(eventCollection, propertyTarget string, percent float64, opts ...queryOpts) Resource

### Sum resource

https://keen.io/docs/api/reference/#sum-resource

    func Sum(eventCollection, propertyTarget string, opts ...queryOpts) Resource

### SelectUnique resource

https://keen.io/docs/api/reference/#select-unique-resource

    func SelectUnique(eventCollection, propertyTarget string, opts ...queryOpts) Resource

### Extraction resource

https://keen.io/docs/api/reference/#extraction-resource

    func Extraction(eventCollection string, opts ...queryOpts) Resource

## License

MIT

