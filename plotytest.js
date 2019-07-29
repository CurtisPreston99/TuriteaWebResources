//this file must be called after the div's that it is editing are other plotly wont be able to find the div's and just throw errors



//the 3 basic data types i think we should support at first

//linegraph
//fist line
var trace1 = {
  x: [1, 2, 3, 4],
  y: [10, 15, 13, 17],
  type: 'scatter',
};
//second line
var trace2 = {
  x: [1, 2, 3, 4],
  y: [16, 5, 11, 9],
  type: 'scatter'
};
//puting the data together
var data = [trace1, trace2];
//add the line graph to the div linegraph
Plotly.newPlot('linegraph', data);


//making bar graph
var data = [
  {
    //names
    x: ['giraffes', 'orangutans', 'monkeys'],
    //height of bars
    y: [20, 14, 23],
    //defining type of graph
    type: 'bar'
  }
];
//adding to dom
Plotly.newPlot('bargraph', data);

//pie chart
var data = [{
  //percentages
  values: [19, 26, 55],
  //data names
  labels: ['Residential', 'Non-Residential', 'Utility'],
  type: 'pie'
}];
//size
var layout = {
  height: 400,
  width: 500
};
//adding to dom
Plotly.newPlot('piechart', data, layout);
