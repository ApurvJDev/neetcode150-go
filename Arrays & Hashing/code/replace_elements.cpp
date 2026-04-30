class Solution {
public:
    vector<int> replaceElements(vector<int>& arr) {
        int n = arr.size();
        vector<int> res(n);
        int maxNum = arr[n - 1];

        for(int i = n - 2; i >= 0; i--) {
            maxNum = max(maxNum, arr[i + 1]);
            res[i] = maxNum;
        }
        res[n - 1] = -1;

        return res;
    }
};
