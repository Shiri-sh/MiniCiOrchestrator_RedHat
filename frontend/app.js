const buildButton = document.getElementById("buildButt")
const repoInput = document.getElementById("repo")
const branchInput = document.getElementById("branch")
const buildsList = document.getElementById("buildsList")
buildButton.addEventListener("click", triggerBuild)
async function triggerBuild() {

  const repo = repoInput.value
  const branch = branchInput.value

  await fetch("http://localhost:8080/build/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ repo, branch })
  })

  alert("Build started")
}

async function fetchBuilds() {
  const response = await fetch("http://localhost:8080/builds")
  const builds = await response.json()
  buildsList.innerHTML = builds.map(build => `
    <tr>
      <td>${build.id}</td>
      <td>${build.repo}</td>
      <td>${build.status}</td>
      <td>${new Date(build.createdAt).toLocaleString()}</td>
    </tr>
    `).join("")
}

setInterval(fetchBuilds, 5000)