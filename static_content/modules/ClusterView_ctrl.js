app.controller("ClusterView_ctrl", function($scope, $window) {


		// Load the initialize the jQuery plugin for the connections.
		// TODO : Once the jq part is dynamic, we'll need to move it here
		
		$window.make_cluster_connections()
});