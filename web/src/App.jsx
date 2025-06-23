import React, { useState, useEffect } from 'react';
import { Send, RefreshCw, List, Eye, Server, CheckCircle, AlertCircle, Loader2 } from 'lucide-react';
import './App.css';

const features = [
  { label: '完整的 A2A 协议支持', enabled: true },
  { label: 'JSON-RPC 2.0 通信', enabled: true },
  { label: 'Server-Sent Events 流式处理', enabled: true },
  { label: '任务生命周期管理', enabled: true },
  { label: 'Agent Card 发现机制', enabled: true },
  { label: '任务查询接口', enabled: true },
  { label: '推送通知（webhooks）', enabled: false },
  { label: '企业级安全', enabled: false },
  { label: '多内容类型（文件、数据等）', enabled: false },
  { label: '动态 Agent 发现/注册', enabled: false },
];

function App() {
  const [inputText, setInputText] = useState('');
  const [isProcessing, setIsProcessing] = useState(false);
  const [agentCard, setAgentCard] = useState(null);
  const [tasks, setTasks] = useState([]);
  const [selectedTask, setSelectedTask] = useState(null);
  const [taskDetail, setTaskDetail] = useState(null);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);
  const [refreshing, setRefreshing] = useState(false);

  // 获取AgentCard信息
  useEffect(() => {
    fetch('/.well-known/agent.json')
      .then(res => res.json())
      .then(setAgentCard)
      .catch(() => setAgentCard(null));
  }, []);

  // 获取所有任务
  const fetchTasks = async () => {
    setRefreshing(true);
    try {
      const res = await fetch('/a2a/tasks');
      if (!res.ok) throw new Error('获取任务列表失败');
      const data = await res.json();
      setTasks(data.sort((a, b) => (b.createdAt || '').localeCompare(a.createdAt || '')));
    } catch (e) {
      setError(e.message);
    } finally {
      setRefreshing(false);
    }
  };

  // 定时刷新任务列表
  useEffect(() => {
    fetchTasks();
    const timer = setInterval(fetchTasks, 3000);
    return () => clearInterval(timer);
  }, []);

  // 获取单个任务详情
  const fetchTaskDetail = async (taskId) => {
    setTaskDetail(null);
    try {
      const res = await fetch(`/a2a/task/${taskId}`);
      if (!res.ok) throw new Error('获取任务详情失败');
      const data = await res.json();
      setTaskDetail(data);
    } catch (e) {
      setError(e.message);
    }
  };

  // 发起新任务
  const sendTask = async () => {
    if (!inputText.trim()) {
      setError('请输入要处理的文本');
      return;
    }
    setIsProcessing(true);
    setError(null);
    setSuccess(null);
    try {
      const messageBody = {
        jsonrpc: "2.0",
        method: "message/send",
        params: {
          message: {
            parts: [{ type: "text", text: inputText }],
            role: "user"
          },
          config: { blocking: false, acceptedOutputModes: ["text", "json"] }
        },
        id: Date.now().toString()
      };
      const response = await fetch('/a2a/server', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(messageBody)
      });
      if (!response.ok) throw new Error('服务器响应错误');
      const result = await response.json();
      if (result.error) throw new Error(result.error.message);
      setSuccess('任务已成功发起！');
      setInputText('');
      fetchTasks();
    } catch (e) {
      setError(e.message);
    } finally {
      setIsProcessing(false);
    }
  };

  // 任务状态颜色
  const statusColor = (status) => {
    switch ((status || '').toLowerCase()) {
      case 'completed': return '#22c55e';
      case 'failed': return '#ef4444';
      case 'working': return '#3b82f6';
      case 'submitted': return '#f59e42';
      default: return '#64748b';
    }
  };

  return (
    <div className="app-container">
      <div className="app-header">
        <h1>🤖 A2A 任务监控演示</h1>
        <p>支持多任务并发发起与实时监控</p>
      </div>
      <div className="dashboard-grid">
        {/* 服务器信息 */}
        <div className="card">
          <h2><Server /> A2A服务器 <span className="version-badge">前端版本：0.01</span></h2>
          {agentCard ? (
            <div className="server-info">
              <div className="info-item"><strong>名称:</strong> {agentCard.name}</div>
              <div className="info-item"><strong>后端版本:</strong> {agentCard.version}</div>
              <div className="info-item"><strong>描述:</strong> {agentCard.description}</div>
              {agentCard.capabilities && Object.entries(agentCard.capabilities).map(([key, value]) => (
                <div className="info-item" key={key}>
                  <strong>{key}:</strong> {value === true ? '✅' : value === false ? '❌' : String(value)}
                </div>
              ))}
            </div>
          ) : <div className="loading">正在获取服务器信息...</div>}
        </div>
        {/* 任务发起区 */}
        <div className="card input-card">
          <h2><Send /> 发起新任务</h2>
          <textarea
            value={inputText}
            onChange={e => setInputText(e.target.value)}
            className="input-field"
            placeholder="请输入要发送的内容..."
            rows={4}
            disabled={isProcessing}
          />
          <button className="btn process-btn" onClick={sendTask} disabled={isProcessing || !inputText.trim()}>
            {isProcessing ? <Loader2 className="spin" size={16} /> : <Send size={16} />}
            {isProcessing ? '处理中...' : '发送'}
          </button>
          {success && <div className="success-message"><CheckCircle size={18} /> {success}</div>}
          {error && <div className="error-message"><AlertCircle size={18} /> {error}</div>}
        </div>
        {/* 任务监控区 */}
        <div className="card tasks-card">
          <h2><List /> 任务列表
            <button className="btn refresh-btn" onClick={fetchTasks} disabled={refreshing} title="刷新">
              {refreshing ? <Loader2 className="spin" size={16} /> : <RefreshCw size={16} />}
            </button>
          </h2>
          <div className="tasks-list">
            {tasks.length === 0 ? <div className="loading">暂无任务</div> : (
              <table className="tasks-table">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>状态</th>
                    <th>创建时间</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                  {tasks.map(task => (
                    <tr key={task.id} className={selectedTask === task.id ? 'selected' : ''}>
                      <td style={{fontFamily:'monospace'}}>{task.id}</td>
                      <td><span style={{color: statusColor(task.status?.state)}}>{task.status?.state || '-'}</span></td>
                      <td>{task.createdAt ? new Date(task.createdAt).toLocaleString() : '-'}</td>
                      <td>
                        <button className="btn view-btn" onClick={() => { setSelectedTask(task.id); fetchTaskDetail(task.id); }}>
                          <Eye size={16} /> 详情
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            )}
          </div>
        </div>
        {/* 任务详情区 */}
        <div className="card task-detail-card">
          <h2><Eye /> 任务详情</h2>
          {!selectedTask && <div style={{opacity:0.7}}>请选择左侧任务查看详情</div>}
          {selectedTask && !taskDetail && <div className="loading">加载中...</div>}
          {taskDetail && (
            <div className="task-detail-content">
              <div><strong>ID:</strong> <span style={{fontFamily:'monospace'}}>{taskDetail.id}</span></div>
              <div><strong>状态:</strong> <span style={{color: statusColor(taskDetail.status?.state)}}>{taskDetail.status?.state}</span></div>
              <div><strong>创建时间:</strong> {taskDetail.createdAt ? new Date(taskDetail.createdAt).toLocaleString() : '-'}</div>
              <div><strong>输入:</strong> <pre className="result-pre">{JSON.stringify(taskDetail.history?.[0]?.parts, null, 2)}</pre></div>
              <div><strong>输出:</strong> <pre className="result-pre">{JSON.stringify(taskDetail.artifacts, null, 2)}</pre></div>
              <div><strong>历史:</strong> <pre className="result-pre">{JSON.stringify(taskDetail.history, null, 2)}</pre></div>
            </div>
          )}
        </div>
        {/* 服务器特性卡片，放在右下角 */}
        <div className="card features-card">
          <h2>服务器特性</h2>
          <ul className="features-list">
            {features.map(f => (
              <li key={f.label} className={f.enabled ? 'feature-enabled' : 'feature-disabled'}>
                {f.enabled ? <span className="feature-icon">✅</span> : <span className="feature-icon">⚪️</span>}
                {f.label}
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
}

export default App; 