defmodule AdventOfCode.Day15 do
  alias :math, as: Math

  @starttime DateTime.utc_now()

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

  def part2_superslow(args, example \\ false) do
    # coords = parseToIntLists(args)
    x_max = if example, do: 20, else: 4_000_000
    y_max = if example, do: 20, else: 4_000_000


    coords = parseToIntLists(args)

    # get coverage
    0..y_max
    |> Enum.reduce_while([],
    fn line, _acc -> covered_tiles = getCoverageForLine(coords, line, 0,  x_max)
      if Integer.mod(line, 1_000) == 0 do
          percentage = Float.round(100*line/y_max, 1)
          time_elapsed = DateTime.diff(DateTime.utc_now(), @starttime)
          forecast = time_elapsed / y_max * line
          hours_elapsed = Float.round(time_elapsed/ 60 / 60 )
          hours_forecast = Float.round(forecast/ 60 / 60)
         IO.puts(IO.ANSI.clear_line() <> "Lines checked: #{line} of #{y_max} (#{percentage} %) Elapsed time (#{hours_elapsed} hours of #{hours_forecast})")
        end

      if MapSet.size(covered_tiles) == x_max  do
        col = MapSet.difference(MapSet.new(0..x_max), getCoverageForLine(coords, line, 0, x_max)) |> MapSet.to_list() |> hd

        {:halt, [line + 4000000 * col]}
      else
        {:cont, []}
      end
    end)
  end

  def parseToIntLists(args) do
    Regex.scan(~r/-?\d+/, args) |> Enum.map(&String.to_integer(hd(&1))) |> Enum.chunk_every(4)
  end

  def outOfRange?([sensorX, sensorY, beaconX, beaconY], xMax, yMax) do
    radius = abs(sensorX - beaconX) + abs(sensorY - beaconY)
    left = sensorX - radius
    right = sensorX + radius
    up = sensorY - radius
    down = sensorY + radius

    Range.disjoint?(0..xMax, left..right) or Range.disjoint?(0..yMax, up..down)
  end

  def coveredRangeInLine([sensorX, sensorY, beaconX, beaconY], line, xMin, xMax) do
    left..right = coveredRangeInLine([sensorX, sensorY, beaconX, beaconY], line)
    if right == -1, do: 0..-1//1, else: max(xMin, left)..min(xMax, right)

  end

  def manhatten_distance([x1,y1], [x2,y2]) do
     x = if x1 >= x2, do: x1 - x2, else: x2 - x1
     y = if y1 >= y2, do: y1 - y2, else: y2 - y1
     x + y
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

  def formatLineTo4Some(line) do
    [sensorX, sensorY, beaconX, beaconY] =
      Regex.scan(~r/-?\d+/, line) |> Enum.map(&String.to_integer(hd(&1)))

    radius = abs(sensorX - beaconX) + abs(sensorY - beaconY)
    left = sensorX - radius
    right = sensorX + radius
    up = sensorY - radius
    down = sensorY + radius
    [left..right, up..down]
  end

  def get_scan_corners(line) do
    [sensorX, sensorY, beaconX, beaconY] =
      Regex.scan(~r/-?\d+/, line) |> Enum.map(&String.to_integer(hd(&1)))

    radius = abs(sensorX - beaconX) + abs(sensorY - beaconY)
    left = [sensorX - radius, sensorY]
    right = [sensorX + radius, sensorY]
    up = [sensorX, sensorY + radius]
    down = [sensorX, sensorY - radius]
    [left, up, right, down]
  end

  def get_limits([[x1, y1],[x2,y2], [x3,y3], [x4,y4]]) do
    [
      Enum.min([x1,x2,x3,x4]),
      Enum.max([y1,y2,y3,y4]),
      Enum.max([x1,x2,x3,x4]),
      Enum.min([y1,y2,y3,y4])
    ] # left, up, right, down
  end

  def cut_out_square(workspace, cutout) do
    [l, u, r, d] = workspace
    [cl, cu, cr, cd] = cutout


    if l > cr or r < cl or u < cd or d > cu do
      [workspace]
    else
      top_right =     [cr+1, u, r, cu+1]   #   I
      middle_right =  [cr+1, cu, r, cd]    #  II
      bottom_right =  [cr+1, cd-1, r, d]   # III

      bottom_middle = [cl, cd-1, cr, d]    # IV

      bottom_left =   [l, cd-1, cl-1, d]  # V
      middle_left =   [l, cu, cl-1, cd]   # VI
      top_left =      [l, u, cl-1, cu+1]  # VII


      top_middle =    [cl, u, cr, cu+1]   # VIII

      [top_right, middle_right, bottom_right, bottom_middle, bottom_left, middle_left, top_left, top_middle]
      |> Enum.map(fn [nl, nu, nr, nd] -> [max(l, nl), min(u, nu), min(r, nr), max(d, nd)] end)
      |> Enum.filter(&valid_cutout?(&1, workspace))
      |> IO.inspect(label: "workspace: #{l}, #{u}, #{r}, #{d}, cutout: #{cl}, #{cu}, #{cr}, #{cd}")
    end
  end

  def valid_cutout?([left, up, right, down], [left_original, up_original, right_original, down_original]) do
    left <= right and down <= up and left_original <= left and right_original >= right and down_original <= down and up_original >= up
  end

  def reduceSearchSpace([otherl..otherr, ou..od], [l..r, u..d]) do
    cond do
      # cuts north east:
      l <= otherl and d >= od and otherl <= r and od >= u ->
        delta = max(r - otherl, d - od)
        [l..(r - delta), (u + delta)..d]

      # cuts north west:
      r >= otherr and d >= od and otherr >= l and od >= u ->
        delta = max(otherr - l, d - od)
        [(l + delta)..r, (u + delta)..d]

      # cuts south east:
      l <= otherl and u <= ou and otherl <= r and ou <= d ->
        delta = max(r - otherl, u - ou)
        [l..(r - delta), u..(d - delta)]

      # cuts south west:
      r >= otherr and u <= ou and otherr >= l and ou <= d ->
        delta = max(otherr - l, u - ou)
        [(l + delta)..r, u..(d - delta)]

      true ->
        [l..r, u..d]
    end
  end

  def tilt_coordinates_by([x, y], rotation) do
    r = Math.sqrt(x**2.0 + y**2.0)
    alpha = get_angle([x, y])

    tilted_x = r * Math.cos(alpha +  rotation) #|> Float.round() |> trunc()
    tilted_y = r * Math.sin(alpha +  rotation) #|> Float.round() |> trunc()
    [tilted_x, tilted_y]
  end

  def get_angle([x, y], degrees \\ false) do
    alpha = cond do
      x == 0 and y == 0 -> 0
      x == 0 and y > 0 -> Math.pi / 2
      x == 0 and y < 0 -> Math.pi/2 * 3

      y >= 0  and x > 0 -> Math.atan(y / x)
      y >= 0  and x < 0 ->  Math.pi  + Math.atan(y / x)
      y < 0 and x < 0 ->  Math.pi + Math.atan(y / x)
      y <= 0 and x > 0 ->  2 * Math.pi + Math.atan(y / x)
      true -> raise "error in get_angle, case not covered"
    end
    if degrees, do: alpha / (2 * Math.pi) * 360, else: alpha
  end

  def tilt_coordinates([x, y]) do
    tilt_coordinates_by([x, y], Math.pi/4) |> Enum.map(&Float.round(&1) |> trunc())
  end

  def untilt_coordinates([x, y]) do
    tilt_coordinates_by([x, y], - Math.pi/4) |> Enum.map(&Float.round(&1) |> trunc())
  end

  def construct_workspace(limit) do
    [_, up] = tilt_coordinates([limit, limit])
    [0, up, up, 0]
  end

  def part2_broken(args, example \\ false) do
    limit = if example, do: 20, else: 4_000_000

    initial_workspaces = [[0,5,5,0]]
    cutouts = [
      [  4, 10, 10, -10],
      [-10,  2, 10, -10],
      [-10, 10, 2, -10],
      [-10, 10, 10, 4]
    ]

    cutouts = args |> String.trim() |> String.split("\n") |> Enum.map(fn line -> get_scan_corners(line) |> Enum.map(&tilt_coordinates/1) end) |> Enum.map(&get_limits/1)
    initial_workspaces = [construct_workspace(limit)]

    Enum.reduce(cutouts, initial_workspaces, fn cutout, workspaces ->
       Enum.map(workspaces, fn workspace -> cut_out_square(workspace, cutout) end)
       |> List.flatten()
       |> Enum.chunk_every(4)
      end)
  end

  def part2(args, example \\ false) do
    coords = parseToIntLists(args)
    limit = if example, do: 20, else: 4_000_000


  end

  def fringe_stream() do

  end

end
