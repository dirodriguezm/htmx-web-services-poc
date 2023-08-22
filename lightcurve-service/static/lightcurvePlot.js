var myChart = echarts.init(document.getElementById('main'));

fetch(`http://localhost:8000/detections/${oid}`).then(
  response => response.json()
).then(data => {
  let g = data.filter(function (item) {
    return item["fid"] == 1
  })
  let r = data.filter(function (item) {
    return item["fid"] == 2
  })
  console.log(g)
  console.log(r)
  var option = {
    tooltip: {},
    legend: {
      data: ['g', 'r']
    },
    xAxis: {
      type: 'value',
      min: function (value) {
        return value.min - 5;
      },
      max: function (value) {
        return value.max + 5;
      }
    },
    yAxis: {
      type: 'value',
      min: function (value) {
        return value.min - 0.5;
      },
      max: function (value) {
        return value.max + 0.5;
      },
      inverse: true
    },
    series: [
      {
        name: 'g',
        type: 'scatter',
        data: g.map(function (item) {
          return [item["mjd"], item["mag"]]
        }),
        color: 'green'
      },
      {
        name: 'r',
        type: 'scatter',
        data: r.map(function (item) {
          return [item["mjd"], item["mag"]]
        }),
        color: 'red'
      }
    ]
  };

  // Display the chart using the configuration items and data just specified.
  myChart.setOption(option);
});
