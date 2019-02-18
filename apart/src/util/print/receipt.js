function getYear(date) {
  return (date / 10000) | 0;
}

function getMonth(date) {
  return ((date / 100) | 0) % 100;
}

function getDate(date) {
  return date % 100;
}

export const html = `
<!doctype html>
<html lang="ja">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>印刷</title>
    <style>
      @page {
        size: A4;
        margin: 0;
      }
      * {
        margin: 0;
        padding: 0;
        font-family: "游ゴシック Medium", "Yu Gothic Medium", "游ゴシック体", YuGothic, sans-serif;
        font-size: 3.5mm;
        font-weight: normal;
        -webkit-print-color-adjust: exact;
      }
      @media print {
        html, body {
          width: 100%;
          height: 100%;
        }
        .receipt-paper {
          width: 100%;
          height: 100%;
        }
      }
      @media screen {
        html,
        body {
          overflow: hidden;
        }
        body {
          background-color: #ddd;
        }
        .paper {
          background-color: #fff;
          width: 210mm;
          height: 297mm;
          margin: 40px;
          border: 1px solid #aaa;
          box-shadow: 5px 5px 5px rgba(0,0,0,0.4);
        }
      }
      .paper {
        overflow: hidden;
        position: relative;
        box-sizing: border-box;
        page-break-after: always;
        height: 297mm;
        width: 210mm;
      }
      .paper > section {
        box-sizing: border-box;
        width: 100%;
        height: calc(100% / 3);
        padding: 15mm 30mm;
      }
      .paper > section:not(:last-child) {
        border-bottom: 1px dashed black;
      }
      .paper > .receipt {
        display: grid;
        grid-template: 'header     header   ' 10mm
                       'tenantName publishAt' 10mm
                       'table      table    ' auto
                       'admin      admin    ' auto
                       /auto       auto;
      }
      .paper > .certificate {
        display: grid;
        grid-template: 'header     header   ' 10mm
                       'tenantName publishAt' 10mm
                       'charge     charge   ' auto
                       'admin      admin    ' auto
                       /auto       auto;
      }
      .publishAt {
        text-align: right;
      }
      .charge-table {
        width: auto;
        border-collapse: collapse;
        border: 1px solid #000;
      }
      .charge-table th {
        font-weight: bold;
      }
      .charge-table thead {
        border-bottom: 1px solid #000;
      }
      .charge-table tbody {
        border-bottom: 1px dotted #000;
      }
      .charge-table td:not(:last-child),
      .charge-table th:not(:last-child) {
        border-right: 1px dotted #000;
      }
      .charge-table th,
      .charge-table td {
        padding: 0.2mm 5mm;
      }
      .receipt .charge {
        text-align: right;
      }
      .certificate .charge {
        text-align: center;
        font-size: 4mm;
      }
      .certificate .charge strong {
        font-size: 5mm;
        font-weight: bold;
      }
      .admin {
        margin-top: 2mm;
        text-align: right;
      }
      .paper > section > h1 {
        font-size: 6mm;
        text-align: center;
      }
      .paper > section > .tenant-name {
        font-size: 4mm;
      }
    </style>
  </head>
  <body></body>
</html>
`;

export function template({
  tenantName,
  rent,
  administrator,
  waterCharge,
  parkingFee,
  commonAreaCharge,
  publishAt
}) {
  const sum =
    Number(rent) +
    Number(waterCharge) +
    Number(parkingFee) +
    Number(commonAreaCharge);
  const receipt = `
  <div class="tenantName" style="grid-area: tenantName;">${tenantName} 様</div>
  <div class="publishAt" style="grid-area: publishAt;">
    発行日:
    ${getYear(publishAt)} 年
    ${getMonth(publishAt)} 月
    ${getDate(publishAt)} 日
  </div>
  <table class="charge-table" style="grid-area: table;">
    <thead>
      <tr>
        <th>項目名</th>
        <th>金額（税込）</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>家賃</td>
        <td class="charge">${(+rent).toLocaleString()} 円</td>
      </tr>
      <tr>
        <td>水道料金</td>
        <td class="charge">${(+waterCharge).toLocaleString()} 円</td>
      </tr>
      <tr>
        <td>駐車場料金</td>
        <td class="charge">${(+parkingFee).toLocaleString()} 円</td>
      </tr>
      <tr>
        <td>共益費</td>
        <td class="charge">${(+commonAreaCharge).toLocaleString()} 円</td>
      </tr>
    </tbody>
    <tfoot>
      <tr>
        <td>計</td>
        <td class="charge">${(+sum).toLocaleString()} 円</td>
      </tr>
    </tfoot>
  </table>
  <div class="admin" style="grid-area: admin;">${administrator}</div>
  `;
  return `
  <article class="paper">
    <section class="receipt">
      <h1 class="header" style="grid-area: header;">家賃請求(控え)</h1>
      ${receipt}
    </section>
    <section class="certificate">
      <h1 style="grid-area: header;">入金証</h1>
      <div class="tenant-name" style="grid-area: tenantName;">
        ${tenantName} 様
      </div>
      <div class="publishAt" style="grid-area: publishAt;">
        発行日: 　　　　 年 　　 月 　　 日
      </div>
      <div class="charge" style="grid-area: charge;">
        金額 <strong>${sum.toLocaleString()} 円</strong> を頂きました。
      </div>
      <div class="admin" style="grid-area: admin;">
        ${administrator}
      </div>
    </section>
    <section class="receipt">
      <h1 style="grid-area: header;">家賃請求</h1>
      ${receipt}
    </section>
  </article>
  `;
}
