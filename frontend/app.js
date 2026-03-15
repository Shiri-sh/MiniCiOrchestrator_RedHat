const buildButton = document.getElementById("buildButt")
const repoInput = document.getElementById("repo")
const branchInput = document.getElementById("branch")
const buildsList = document.getElementById("buildsList")
buildButton.addEventListener("click", (event) => triggerBuild(event))
async function triggerBuild(event) {
  event.preventDefault()
  const repo = repoInput.value
  const branch = branchInput.value

  try {
    const response = await fetch("http://localhost:8080/build/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ repo, branch })
    })
  }
  catch (error) {
    console.log("Error triggering build:", error.message)
    return
  }

}

async function fetchBuilds() {
  try {
    console.log("Fetching builds...")
    const response = await fetch("http://localhost:8080/builds")
    const builds = await response.json()
    console.log("Fetched builds:", builds)
    buildsList.innerHTML = builds.map(build => `
    <tr>
      <td>${build.id}</td>
      <td>${build.repo}</td>
      <td>${build.status}</td>
      <td>${new Date(build.created_at).toLocaleString()}</td>
    </tr>
    `).join("")
  }
  catch (error) {
    console.log("Error fetching builds:", error.message)
  }
}

//setInterval(fetchBuilds, 1000)