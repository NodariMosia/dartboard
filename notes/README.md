# Dartboard Dimensions and Darts Scoring System

## Dartboard Diagram

![dartboard](dartboard.png)

## General Dartboard Dimensions

| Measurement                    | Value (mm) |
| ------------------------------ | ---------- |
| Width (Full Board Diameter)    | 451        |
| Width (Scoring Part Diameter)  | 340        |
| Double Bullseye Diameter       | 12.7       |
| Single Bullseye Diameter       | 32         |
| Center to Outside Treble       | 107        |
| Center to Outside Double       | 170        |
| Width of Double & Treble Rings | 8          |

## Dartboard Zone Dimensions

Start of the zone is the distance from the center of the dartboard to the
starting of the zone. It represents the radius of the zone's starting circle.

End of the zone is the distance from the center of the dartboard to the
end of the zone. It represents the radius of the zone's ending circle.

Width of the zone is the distance between the start and end of the zone. It
represents the width of the ring between zone's starting and ending circles.

| Zone            | Start (mm) | End (mm) | Width (mm) |
| --------------- | ---------- | -------- | ---------- |
| Double Bullseye | 0          | 6.35     | 6.35       |
| Single Bullseye | 6.35       | 16       | 9.65       |
| Inner Single    | 16         | 99       | 83         |
| Treble          | 99         | 107      | 8          |
| Outer Single    | 107        | 162      | 55         |
| Double          | 162        | 170      | 8          |

## Dartboard Sectors

The dartboard has 20 equally sized sectors (18 degrees). Each sector has a
unique score from 1 through 20. The sectors are ordered in the following fashion
(starting from the top center, going clockwise):

20, 1, 18, 4, 13, 6, 10, 15, 2, 17, 3, 19, 7, 16, 8, 11, 14, 9, 12, 5.

## Darts Scoring System

### Score is calculated as follows:

- In case of double bullseye: score = 50;
- In case of single bullseye: score = 25;
- In other cases: score = sector score \* zone multiplier.

### Zone Score / Multiplier Table

| Zone            | Zone Score | Zone Multiplier |
| --------------- | ---------- | --------------- |
| Double Bullseye | 50         | N/A             |
| Single Bullseye | 25         | N/A             |
| Inner Single    | N/A        | 1               |
| Treble          | N/A        | 3               |
| Outer Single    | N/A        | 1               |
| Double          | N/A        | 2               |

## References

- https://en.wikipedia.org/wiki/Darts
- https://www.dimensions.com/element/dartboard
