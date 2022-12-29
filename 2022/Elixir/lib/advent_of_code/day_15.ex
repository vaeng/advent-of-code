defmodule AdventOfCode.Day15 do

  def part1(args) do
    coords = parseToIntLists(args)
    line = 2_000_000
    # get Beacons in the line
    beaconsInLine =
      coords
      |> Enum.filter(fn [_sensorX, _sensorY, _beaconX, beaconY] -> beaconY == line end)
      |> Enum.uniq_by(fn [_sensorX, _sensorY, beaconX, _beaconY] -> beaconX end)
      |> Enum.count()

    # get coverage
    coverage =
      getCoverageForLine(coords, line, -40_000_000, 40_000_000)
      |> MapSet.size()

    coverage - beaconsInLine
  end

  def getCoverageForLine(coords, line, xMin, xMax) do
    coords
    |> Enum.map(&coveredRangeInLine(&1, line, xMin, xMax))
    |> Enum.reduce(MapSet.new(), fn rng, acc -> MapSet.union(acc, MapSet.new(rng)) end)
  end

  def coveredRangeInLine([sensorX, sensorY, beaconX, beaconY], line, xMin, xMax) do
    left..right = coveredRangeInLine([sensorX, sensorY, beaconX, beaconY], line)
    if right == -1, do: 0..-1//1, else: max(xMin, left)..min(xMax, right)
  end

  def coveredRangeInLine([sensorX, sensorY, beaconX, beaconY], line) do
    radius = manhatten_distance([sensorX, sensorY], [beaconX, beaconY])
    distanceToLine = abs(line - sensorY)
    left = sensorX - radius + distanceToLine
    right = sensorX + radius - distanceToLine

    if distanceToLine <= radius do
      left..right
    else
      # empty range
      0..-1//1
    end
  end

  def parseToIntLists(args) do
    Regex.scan(~r/-?\d+/, args) |> Enum.map(&String.to_integer(hd(&1))) |> Enum.chunk_every(4)
  end

  def manhatten_distance([x1,y1], [x2,y2]) do
     x = if x1 >= x2, do: x1 - x2, else: x2 - x1
     y = if y1 >= y2, do: y1 - y2, else: y2 - y1
     x + y
  end

  def part2(args, example \\ false) do
    coords = parseToIntLists(args)
    limit = if example, do: 20, else: 4_000_000
    initial_fringes = Stream.flat_map(coords, fn coord -> fringe_stream(coord, limit) end)

    [col, line] = Enum.reduce(coords, initial_fringes, fn coord, fringes -> Stream.reject(fringes, &is_covered?(coord, &1)) end) |> Stream.uniq() |> Enum.to_list() |> hd
    line + 4000000 * col
  end

  def fringe_stream([sensorX, sensorY, beaconX, beaconY], limit) do
    radius = manhatten_distance([sensorX, sensorY], [beaconX, beaconY])
    left = sensorX - radius
    right = sensorX + radius
    up = sensorY - radius
    down = sensorY + radius

    up_to_right = make_streamed_line([sensorX, up-1], [right+1, sensorY])
    right_to_down = make_streamed_line([right+1, sensorY], [sensorX, down+1])
    down_to_left = make_streamed_line([sensorX, down+1], [left-1, sensorY])
    left_to_up = make_streamed_line([left-1, sensorY], [sensorX, up-1])

    Stream.concat([up_to_right, right_to_down, down_to_left, left_to_up])
    |> Stream.filter(fn [x,y] -> x >= 0 and y >= 0 and y <= limit and x <= limit end)
  end

  def make_streamed_line([start_x, start_y], [exclusive_end_x, exclusive_end_y]) do
    Stream.zip_with(start_x..(exclusive_end_x-1), start_y..(exclusive_end_y-1), fn  x,y -> [x,y] end)
  end

  def is_covered?([sensorX, sensorY, beaconX, beaconY], [x,y]) do
    radius = manhatten_distance([sensorX, sensorY], [beaconX, beaconY])

    distanceToLine = abs(y - sensorY)
    left = sensorX - radius + distanceToLine
    right = sensorX + radius - distanceToLine

    x >= left and x <= right
  end

end
