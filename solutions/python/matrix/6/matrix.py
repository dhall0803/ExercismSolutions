"""Module providing a class (Matrix) which is a representation of a matrix of numbers"""

class Matrix:
    """Represents a matrix of numbers"""
    def __init__(self, matrix_string) -> None:
        self.rows = self.rows_from_string(matrix_string)

    @staticmethod
    def rows_from_string(matrix_string) -> list[list[int]]:
        """Takes a string representing a matrix of numbers, where a new row is indicated with a newline, and returns as list of lists of integers where each list is a row in the matrix"""
        rows = []
        for line in matrix_string.splitlines():
            rows.append([int(digit) for digit in line.split()])
        return rows
        
    def row(self, index) -> list[int]:
        """Returns the row at a 1 indexed index"""
        return self.rows[index - 1]

    def column(self, index) -> list[int]:
        """Returns the column at the given one indexed index"""
        return [row[index - 1] for row in self.rows]
