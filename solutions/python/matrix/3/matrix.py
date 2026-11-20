class Matrix:
    def __init__(self, matrix_string):
        self.rows = self.rows_from_string(matrix_string)

    def rows_from_string(self, matrix_string):
        rows = []
        for line in matrix_string.splitlines():
            digits = line.split()
            rows.append([ int(digit) for digit in digits ])
        return rows
        
    def row(self, index):
        return self.rows[index - 1]

    def column(self, index):
        return [row[index - 1] for row in self.rows]
