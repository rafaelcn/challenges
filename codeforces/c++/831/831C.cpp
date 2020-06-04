// Author: Rafael Campos Nunes
// Date: 22/07/2017


// This solution still doesn't work because of its time complexity.

#pragma GCC optimize ("Ofast")

#include <bits/stdc++.h>

using ll = long long int;

int
main() {
    using namespace std;

    int k, n;
    
    cin >> k >> n;

    int judge_values[k];
    int heard_values[n];
    
    for (int i = 0; i < k; i++) {
	cin >> judge_values[i];
    }

    for (int i = 0; i < n; i++) {
	cin >> heard_values[i];
    }

    // set used to test intersection
    set<ll> intersection;
    set<ll> intersection1;

    vector<ll> result;
    
    // O(n*k)
    for (int i = 0; i < n; i++) {
        // last sum
        ll ls = 0;
        
    	for (int j = 0; j < k; j++) {
            ls = ls + judge_values[j];
            ll value = heard_values[i] - ls;
	    
            if (i == 0) {
		// O(logn)
                intersection.insert(value);
            } else {
                intersection1.insert(value);
            }
    	}
	
	if (n == 1) {
	    for (const auto &e : intersection) {
		result.push_back(e);
	    }
	} else {
	    if (i != 0) {
		// at most 2*(N1+N2-1)
		set_intersection(intersection.begin(), intersection.end(),
				 intersection1.begin(), intersection1.end(),
				 std::back_inserter(result));

		// clearing the old vector, with complexity of O(n)
		intersection.clear();

		// Also O(n)
		for (const auto &e : result) {
		    intersection.insert(e);
		}
	    }
	}
    }

    // Total complexity: O(n*k) + 2*(n-1)*O(n) + (n-1)*O(2*(N1+N2-1))...

    // I'm not sure why but using %zu or %lld (which might seem correct)
    // doesn't work on CodeForces, maybe because it's windows.
    printf("%lu", intersection.size());

    return 0;
}
