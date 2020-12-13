Day 12
======

## Part 1
The problem is pretty easy. The interesting part here is how to implement moving
to different directions.

We can think of east/west, north/south coordinates as a point of 2d plane. If we
are moving to east it means we are increasing our X coordinates, if we are
moving to west we are decreasing our X coordinates. The same with north and
easy, but with Y coordinate. So, for representing directions on the 2d plane we
can make an mapping of changes on every direction. For example, the change for
East direction is +1, 0 which means we have to increase our X coordinate and
don't change Y coordinate.

The rotation on the plane can be encoded as a circular list of directions. E.g.
we are facing to East rotating right means to face to South, after next right
rotate we'll face West and then North and then again East. If we rotate left the
order changes.


## Part 2

The hard part here is how to encode waypoint rotation. I used rotation matrix
for that and just multiplied the matrix with the current coordinate vector.
