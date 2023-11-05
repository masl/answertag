// wait for the DOM to load and register our event listeners
// additionally run the function once on page load
document.addEventListener("DOMContentLoaded", () => {
  drawCloud();

  const mainElement = document.getElementsByTagName("main")[0];
  mainElement.addEventListener("htmx:wsAfterMessage", function () {
    drawCloud();
  });
});

// draw the tag cloud using d3.js
function drawCloud() {
  var tags = document.querySelectorAll("#tag-list > li");
  const wordData = Array.from(tags).map((tag) => {
    return {
      text: tag.dataset.name,
      size: tag.dataset.fontSize,
    };
  });

  // set the dimensions and margins of the graph
  const margin = { top: 10, right: 10, bottom: 10, left: 10 },
    width = 450 - margin.left - margin.right,
    height = 450 - margin.top - margin.bottom;

  // append the svg object to the body of the page
  const svg = d3
    .select("#tagcloud")
    .append("svg")
    .attr("width", width + margin.left + margin.right)
    .attr("height", height + margin.top + margin.bottom)
    .append("g")
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

  // Constructs a new cloud layout instance. It run an algorithm to find the position of words that suits your requirements
  const layout = d3.layout
    .cloud()
    .size([width, height])
    .words(wordData)
    .rotate(0)
    .padding(10)
    .fontSize((d) => d.size)
    .on("end", draw);
  layout.start();

  // This function takes the output of 'layout' above and draw the words
  // Better not to touch it. To change parameters, play with the 'layout' variable above
  function draw(words) {
    svg
      .append("g")
      .attr(
        "transform",
        "translate(" + layout.size()[0] / 2 + "," + layout.size()[1] / 2 + ")"
      )
      .selectAll("text")
      .data(words)
      .enter()
      .append("text")
      .style("font-size", function (d) {
        return d.size + "px";
      })
      .attr("text-anchor", "middle")
      .attr("transform", function (d) {
        return "translate(" + [d.x, d.y] + ")rotate(" + d.rotate + ")";
      })
      .text(function (d) {
        return d.text;
      });
  }
}
