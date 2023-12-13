defmodule Aoc23.D12 do
  require Logger
  def run() do
    case  File.read("input/d12/input") do
      {:ok, file} ->
        String.split(file,"\n")
        |> Enum.map(&parseLine/1)
        |> String.sum()

      {:error, reason} ->
        Logger.error(reason)
    end
  end

  def parseLine(line) do
    String.
  end
end
