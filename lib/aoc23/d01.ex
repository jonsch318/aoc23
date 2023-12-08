defmodule Aoc23.D01 do
  require Logger
  def run() do
    case  File.read("input/d01/input.txt") do
      {:ok, file} ->
        file
        |> String.split("\n")
        |> Enum.map(&parseLine/1)
        |> Enum.map(&{List.first(&1), List.last(&1)})
        |> Enum.map(fn
          {nil, nil} -> 0
          {nil, r} -> r
          {l,r} -> (l * 10) + r end)
        |> Enum.sum()
        |> IO.puts
      {:error, reason} ->
        Logger.error(reason)
    end
  end

  def parseLine(line) do
    line
    |> String.graphemes()
    |> Enum.map(fn <<c>> -> c end) # to binary
    |> Enum.filter(&(&1 in ?1..?9)) # e.g. Integer.parse
    |> Enum.map(&(&1 - ?0)) # remove 0

  end
end
