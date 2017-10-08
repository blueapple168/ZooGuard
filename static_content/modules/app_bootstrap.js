var app = angular.module("zg", ['ngRoute']);


 // configure our routes
app.config(function($routeProvider, $locationProvider) {

    $routeProvider

        // route for the home page
        .when('/', {

            templateUrl : 'views/home.html',
            controller  :  'Home_ctrl'
        })
        .when('/cluster_view', {

            templateUrl : 'views/cluster_view.html',
            controller  :  'ClusterView_ctrl'
        })
        .when('/component_view', {
            
            templateUrl : 'views/component_list.html',
            controller  :  'ComponentList_ctrl'
        })

        .when('/component_detail/:component_name', {
            
            templateUrl : 'views/component_detail.html',
            controller  :  'ComponentDetail_ctrl'
        })



});

// Will be called on all the pages
app.run(function($rootScope) {

    $rootScope.$on('$routeChangeStart', function(next, current) { 
            

            // The damn thing is putting its shit outside the BODY tags... 
            window.jQuery(".connection").remove();
     });

})

