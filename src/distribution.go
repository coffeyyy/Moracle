import (
	"fmt"
	"math"
)
// sample pricing data (ETH - 03/24/2025 - Around 9:30 PM EST), size = 30
prices = [30]float{2082.3327208592864, 2081.4219170617016, 2082.1769288658647, 2081.6392055267524, 2082.0958889816566, 2081.385814173955, 2081.8784905497478, 2081.7225096445864, 2081.619691914775, 2081.6453400564164, 2081.903964869011, 2081.8158877018373, 2081.87691761155, 2081.5832583273, 2081.559575088938, 2081.8743546467294, 2082.3057356005825, 2082.3528199265743, 2081.9101853250004, 2081.7457861889598, 2081.5419083716192, 2081.8599222167322, 2081.47807590254, 2082.171883125856, 2082.1469488090866, 2082.0246213156315, 2081.9758774578904, 2082.1928023583487, 2081.6543804741464, 2081.7407762419816}


func create_distribution() {

}

// The implied population mean is the average price between Pyth, Chainlink, Redstone, and Chronicle
func calculate_t()

func calculate_standard_deviation(float mean, int size, float [30]prices) {
	sum_square_diffs = 0
	for _, price := range prices {
		sum_square_diffs += ((prices[i] - mean) ** 2)
	}
	return math.Sqrt(sum_square_diffs/size-1)
}

func calculate_mean(arr [30]prices, int size) {
	sum = 0
	for _, price := range prices {
		sum += price
	}
	return sum/size
}