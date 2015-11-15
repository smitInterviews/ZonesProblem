Zipcode to Zones Problem
=============================

The server is deployed at 52.26.157.247 running under the port 8080. A GET request can be made based on this template: 

``` curl <public_ip>:<port>/zone/<origin_zip_code>/distancematrix?dest_code=<destination_zip_code> ```
The response is the standard http codes (404 for not found, 400 for bad request, 500 for internal error, 200 for successfully processed request). The response is of type json with key being "Message" and value being "zone <x>" where x is the relevant zone associated with the given destination zip code. 

All source code is under server (logic for csv parsing, server init, and zone searching based on zip codes). All tests are under /tests.  


