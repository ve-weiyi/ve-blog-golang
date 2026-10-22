<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{title}</title>
</head>
<body style="margin:0;padding:0;background:#f4f6fb;font-family:-apple-system,BlinkMacSystemFont,Segoe UI,PingFang SC,Microsoft YaHei,sans-serif;">
  <table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="background:#f4f6fb;padding:32px 16px;">
    <tr>
      <td align="center">
        <table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="max-width:560px;background:#ffffff;border-radius:16px;overflow:hidden;box-shadow:0 8px 30px rgba(15,23,42,0.08);">
          <tr>
            <td style="background:#4f46e5;background:linear-gradient(135deg,#4f46e5,#7c3aed);padding:28px 32px;color:#ffffff;">
              <div style="font-size:20px;font-weight:700;letter-spacing:1px;">博客</div>
              <div style="margin-top:6px;font-size:13px;opacity:0.85;">安全验证邮件</div>
            </td>
          </tr>
          <tr>
            <td style="padding:32px;">
              <h1 style="margin:0 0 12px;font-size:20px;color:#111827;">{title}</h1>
              <p style="margin:0 0 24px;font-size:14px;line-height:1.8;color:#4b5563;">{intro}</p>
              <div style="background:#f5f3ff;border:1px solid #ddd6fe;border-radius:12px;padding:20px;text-align:center;">
                <div style="font-size:12px;color:#6b7280;margin-bottom:8px;">你的验证码</div>
                <div style="font-size:30px;font-weight:700;letter-spacing:8px;color:#4f46e5;font-family:SFMono-Regular,Consolas,monospace;">{code}</div>
              </div>
              <p style="margin:20px 0 0;font-size:13px;color:#9ca3af;">验证码 {time} 分钟内有效。若非本人操作，请忽略此邮件。</p>
            </td>
          </tr>
          <tr>
            <td style="padding:20px 32px;background:#f9fafb;border-top:1px solid #f0f0f5;">
              <p style="margin:0;font-size:12px;color:#9ca3af;">此邮件由系统自动发送，请勿直接回复。</p>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>
</html>
