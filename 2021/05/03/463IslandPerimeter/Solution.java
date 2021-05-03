class Solution {
    public int islandPerimeter(int[][] grid) {
        int area = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] == 1) {
                    if (i-1 < 0 || grid[i-1][j] == 0) {
                        area++;
                    }
                    if (i+1 >= grid.length || grid[i+1][j] == 0) {
                        area++;
                    }
                    if (j-1 < 0 || grid[i][j-1] == 0) {
                        area++;
                    }
                    if (j+1 >= grid[0].length || grid[i][j+1] == 0) {
                        area++;
                    }
                }
            }
        }
        return area;
    }
}