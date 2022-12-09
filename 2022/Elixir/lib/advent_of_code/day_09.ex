defmodule AdventOfCode.Day09 do
  def part1(args) do
    args
    |> parseInput()
    |> headPositions()
    # 1
    |> tailPositions()
    |> MapSet.new()
    |> MapSet.size()
  end

  def part2(args) do
    args
    |> parseInput()
    |> headPositions()
    |> IO.inspect()
    # 1
    |> nextKnotPositions()

    # 2
    |> nextKnotPositions()
    # 3
    |> nextKnotPositions()
    # 4
    |> nextKnotPositions()
    # 5
    |> nextKnotPositions()
    # 6
    |> nextKnotPositions()
    # 7
    |> nextKnotPositions()
    # 8
    |> nextKnotPositions()
    # 9
    |> nextKnotPositions()
    |> MapSet.new()
    |> MapSet.size()
  end

  def draw(positions) do
  end

  def headPositions(instructions) do
    instructions
    |> String.codepoints()
    |> Enum.scan([0, 0], fn x, acc -> AdventOfCode.Day09.updateHead(acc, x) end)
  end

  def tailPositions(previousKnotPositions) do
    previousKnotPositions
    |> Enum.scan([0, 0], fn head, acc -> AdventOfCode.Day09.updateTail(head, acc) end)
  end

  def nextKnotPositions(previousKnotPositions) do
    previousKnotPositions
    |> Enum.scan([0, 0], fn head, acc -> AdventOfCode.Day09.updateTailPart2(head, acc) end)
  end

  def parseInput(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(
      &(String.split(&1, " ")
        |> then(fn [dir, num] -> String.duplicate(dir, String.to_integer(num)) end))
    )
    |> Enum.reduce("", fn x, y -> y <> x end)
  end

  def updateHead([x, y], "R"), do: [x + 1, y]
  def updateHead([x, y], "L"), do: [x - 1, y]
  def updateHead([x, y], "U"), do: [x, y + 1]
  def updateHead([x, y], "D"), do: [x, y - 1]

  def updateTail([xh, yh], [xt, yt]) do
    cond do
      # same row, head right
      yh == yt and xh - xt > 1 -> [xt + 1, yt]
      # same row head left
      yh == yt and xh - xt < -1 -> [xt - 1, yt]
      # same line head up
      xh == xt and yh - yt > 1 -> [xt, yt + 1]
      # same line head down
      xh == xt and yh - yt < -1 -> [xt, yt - 1]
      # head is two up and right
      yh > yt + 1 and xh != xt -> [xh, yt + 1]
      # head is two down and left
      yh < yt - 1 and xh != xt -> [xh, yt - 1]
      xh > xt + 1 and yh != yt -> [xt + 1, yh]
      xh < xt - 1 and yh != yt -> [xt - 1, yh]
      abs(xh - xt) <= 1 and abs(yh - yt) <= 1 -> [xt, yt]
      true -> IO.inspect("currently: head: [#{xh},#{yh}] tail: [#{xt},#{yt}]\n")
    end
  end

  def updateTailPart2(h, t) do
    cond do
      h == getN(t) -> t
      h == getNN(t) -> getN(t)
      h == getS(t) -> t
      h == getSS(t) -> getS(t)
      h == getW(t) -> t
      h == getWW(t) -> getW(t)
      h == getE(t) -> t
      h == getEE(t) -> getE(t)

      h == getSE(t) -> t
      h == getSSE(t) -> getSE(t)
      h == getSEE(t) -> getSE(t)
      h == getSESE(t) -> getSE(t)

      h == getNE(t) -> t
      h == getNNE(t) -> getNE(t)
      h == getNEE(t) -> getNE(t)
      h == getNENE(t) -> getNE(t)

      h == getSW(t) -> t
      h == getSSW(t) -> getSW(t)
      h == getSWW(t) -> getSW(t)
      h == getSWSW(t) -> getSW(t)

      h == getNW(t) -> t
      h == getNNW(t) -> getNW(t)
      h == getNWW(t) -> getNW(t)
      h == getNWNW(t) -> getNW(t)
      true -> IO.inspect("currently: head: [#{h}] tail: [#{t}]\n")
    end
  end

  def getN([x, y]), do: [x, y+1]
  def getNN([x, y]), do: [x, y+2]
  def getS([x, y]), do: [x, y-1]
  def getSS([x, y]), do: [x, y-2]
  def getW([x, y]), do: [x-1, y]
  def getWW([x, y]), do: [x-2,y]
  def getE([x, y]), do: [x+1, y]
  def getEE([x, y]), do: [x+2,y]

  def getSE([x, y]), do: [x+1,y-1]
  def getSSE([x, y]), do: [x+1,y-2]
  def getSEE([x, y]), do: [x+2,y-1]
  def getSESE([x, y]), do: [x+2,y-2]

  def getNE([x, y]), do: [x+1,y+1]
  def getNNE([x, y]), do: [x+1,y+2]
  def getNEE([x, y]), do: [x+2,y+1]
  def getNENE([x, y]), do: [x+2,y+2]

  def getSW([x, y]), do: [x-1,y-1]
  def getSSW([x, y]), do: [x-1,y-2]
  def getSWW([x, y]), do: [x-2,y-1]
  def getSWSW([x, y]), do: [x-2,y-2]

  def getNW([x, y]), do: [x-1,y+1]
  def getNNW([x, y]), do: [x-1,y+2]
  def getNWW([x, y]), do: [x-2,y+1]
  def getNWNW([x, y]), do: [x-2,y+2]

end
