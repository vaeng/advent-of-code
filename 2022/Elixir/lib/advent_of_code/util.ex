defmodule AdventOfCode.Util do
  @moduledoc """
  A utility module that implements functions for performing operations that
  might be used by different problem solutions.


  This module is using named directions in a 2D Space.
  The center X be [0,0]. eg. NN is [0,2].
  |      |     |    |     |      |
  |:----:|:---:|:---:|:---:|:----:|
  | NWNW | NNW | NN | NNE | NENE |
  | NWW  | NW  | N  | NE  | NEE  |
  | WW   | W   | X  | E   | EE   |
  | SWW  | SW  | S  | SE  | SEE  |
  | SWSW | SSW | SS | SSE | SESE |


  """

  @doc """
  Convert a newline- separated String of numbers into
  a list of integers.

  Examples

  iex> AdventOfCode.Util.stringToIntArray("1\n2\n334\n4")
  [1,2,334,4]
  """
  @spec stringToIntArray(String.t()) :: list(integer)
  def stringToIntArray(xs),
    do:
      xs
      |> String.strip()
      |> String.split("\n")
      |> Enum.map(&String.to_integer(&1))

  def getN([x, y]), do: [x, y + 1]
  def getNN([x, y]), do: [x, y + 2]
  def getS([x, y]), do: [x, y - 1]
  def getSS([x, y]), do: [x, y - 2]
  def getW([x, y]), do: [x - 1, y]
  def getWW([x, y]), do: [x - 2, y]
  def getE([x, y]), do: [x + 1, y]
  def getEE([x, y]), do: [x + 2, y]

  def getSE([x, y]), do: [x + 1, y - 1]
  def getSSE([x, y]), do: [x + 1, y - 2]
  def getSEE([x, y]), do: [x + 2, y - 1]
  def getSESE([x, y]), do: [x + 2, y - 2]

  def getNE([x, y]), do: [x + 1, y + 1]
  def getNNE([x, y]), do: [x + 1, y + 2]
  def getNEE([x, y]), do: [x + 2, y + 1]
  def getNENE([x, y]), do: [x + 2, y + 2]

  def getSW([x, y]), do: [x - 1, y - 1]
  def getSSW([x, y]), do: [x - 1, y - 2]
  def getSWW([x, y]), do: [x - 2, y - 1]
  def getSWSW([x, y]), do: [x - 2, y - 2]

  def getNW([x, y]), do: [x - 1, y + 1]
  def getNNW([x, y]), do: [x - 1, y + 2]
  def getNWW([x, y]), do: [x - 2, y + 1]
  def getNWNW([x, y]), do: [x - 2, y + 2]

  def transpose(rows) do
    rows
    |> List.zip()
    |> Enum.map(&Tuple.to_list/1)
  end

  @type direction :: :north | :east | :west | :south | :up | :down

  @spec getNext([...], [...] | direction) :: [...]
  def getNext([x, y, z], :north), do: [x, y + 1, z]
  def getNext([x, y, z], :east), do: [x + 1, y, z]
  def getNext([x, y, z], :south), do: [x, y - 1, z]
  def getNext([x, y, z], :west), do: [x - 1, y, z]
  def getNext([x, y, z], :up), do: [x, y, z + 1]
  def getNext([x, y, z], :down), do: [x, y, z - 1]
  def getNext([x, y, z], directions), do: Enum.reduce(directions, [x, y, z], &getNext(&2, &1))
end
