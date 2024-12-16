### **Checklist for Building the CLI App with Cobra**

---

#### **1. Initial Setup**
- [ ] Install Go and Cobra CLI (`go install github.com/spf13/cobra-cli`).
- [ ] Create the project using `cobra-cli init`.
- [ ] Test the base command (`go run main.go`).

---

#### **2. CLI Structure**
- [ ] Add commands to the CLI:
  - [ ] `monitor`: To monitor directories.
  - [ ] `report`: To generate reports.
  - [ ] `setup`: To configure directories and preferences.
- [ ] Define flags for the commands:
  - [ ] `--directory` for the `monitor` command to select the directory.
  - [ ] `--format` for the `report` command to choose the output format (e.g., table, JSON).

---

#### **3. Implementation**
- **File Monitoring**:
  - [ ] Use `fsnotify` to monitor changes in files.
  - [ ] Log the following:
    - File names of modified files.
    - Added/removed lines.
    - File extensions.
- **Git Data Capture**:
  - [ ] Execute Git commands (`git log`, `git diff`) to capture:
    - Commits made.
    - Commit messages.
    - Branches used.
- **Data Storage**:
  - [ ] Save all captured data to a JSON file in the configured directory.

---

#### **4. Reports**
- [ ] Implement the `report` command to:
  - [ ] Read data from the JSON file.
  - [ ] Generate reports in:
    - Table format (for terminal visualization).
    - JSON format (for export or detailed analysis).

---

#### **5. Configuration**
- [ ] Implement the `setup` command to:
  - [ ] Set the default directory for monitoring.
  - [ ] Save preferences to a JSON configuration file (`config.json`).

---

#### **6. Testing**
- [ ] Test the CLI commands:
  - [ ] Test `monitor` with different directories.
  - [ ] Test `report` in different output formats.
  - [ ] Test `setup` and verify the creation of the configuration file.
- [ ] Validate expected behavior for errors (e.g., non-existent directories, invalid JSON).

