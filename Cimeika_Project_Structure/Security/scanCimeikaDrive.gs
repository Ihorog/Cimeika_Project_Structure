function scanCimeikaDrive() {
  const folderId = '1vKLDCP2tX44inRkdFDSG6riOP5YmC2HU';
  const folder = DriveApp.getFolderById(folderId);
  const files = folder.getFiles();
  let result = [];

  while (files.hasNext()) {
    const file = files.next();
    const content = file.getBlob().getDataAsString();
    if (/apiKey|token|secret|authDomain|projectId|firebase/i.test(content)) {
      result.push({
        name: file.getName(),
        url: file.getUrl(),
        type: file.getMimeType()
      });
    }
  }
  Logger.log(JSON.stringify(result, null, 2));
}
