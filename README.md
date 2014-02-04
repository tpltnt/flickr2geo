flickr2geo
==========

A simple scraper to extract geo-data from flickr, nothing fancy. If you are up to a sophisticated use, check the [flickr API](https://secure.flickr.com/services/api/).
The output labels latitude and longiture explicitly and has a final [awk](https://en.wikipedia.org/wiki/Awk)-friendly line. The latitude for example can be extracted with ```./flickr2geo flickr-url | grep awk | awk '{print $3}'```. Released under [AGPLv3 (or later versions)](https://www.gnu.org/licenses/agpl-3.0.html).
