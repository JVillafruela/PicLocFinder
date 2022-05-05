# PicLocFinder

Find geotagged photos according to the location where they were taken

Target location can be :
 - [x] a bounding box
 - [x] a circle (v0.2)
 - [X] a list of points in a GeoJSON file (v0.3)


```
NAME:                                                                                                          
   plf - Picture Location Finder                                                                               
                                                                                                               
   Find geotagged photos according to the location where they were taken                                       
                                                                                                               
USAGE:                                                                                                         
   plf [global options] [arguments...]                                                                         
                                                                                                               
GLOBAL OPTIONS:                                                                                           
   --bbox value                    bounding box ("lon1,lat1,lon2,lat2")                                   
   --latitude value, --lat value   latitude (WGS84 [-90,+90]) (default: 0)                                
   --longitude value, --lon value  longitude (WGS84 [-180,+180]) (default: 0)                             
   --radius value                  radius (in meters) (default: 0)                                        
   --file value, -f value          geoJSON file                                                           
   --help, -h                      show help                                                              
   --version, -v                   print the version                                                      
                                                                                                          
                                                                                                          
WEBSITE: https://github.com/JVillafruela/PicLocFinder                                                     
                                                                                                          
EXAMPLES:                                                                                                 
   plf --bbox="5.68678,45.08596,5.68979,45.08778" E:\OSM\gps\2022\2022-04-16 E:\OSM\gps\2022\2022-04-22   
                                                                                                          
   plf --latitude=45.087 --longitude=5.688 --radius=20  E:\OSM\gps\2022                                   
                                                                                                          
   plf --file C:\OSM\guideposts\guidepost-create.geojson  --radius 20 C:\OSM\Traces\2022\                 
                                                                                                          
```